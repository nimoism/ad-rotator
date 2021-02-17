package grpc //nolint:dupl

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
)

func TestCRUDUserGroup(t *testing.T) {
	apiHost := os.Getenv("ADR_TEST_API_HOST")
	if apiHost == "" {
		t.Skipf("Skip integration test")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, apiHost, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	client := pb.NewUserGroupsClient(conn)
	resAll, err := client.UserGroups(ctx, &pb.AllUGsRequest{})
	require.NoError(t, err)
	require.Empty(t, resAll.Ugs)

	resCreate, err := client.CreateUserGroup(ctx, &pb.CreateUGRequest{
		Ug: &pb.CreateUGRequest_NewUserGroup{Name: "usre_gourp0"},
	})
	require.NoError(t, err)
	pbUG0 := resCreate.Ug
	require.Greater(t, int(pbUG0.Id), 0)
	require.Equal(t, "usre_gourp0", pbUG0.Name)

	resAll, err = client.UserGroups(ctx, &pb.AllUGsRequest{})
	require.NotNil(t, resAll)
	require.NoError(t, err)
	require.Len(t, resAll.Ugs, 1)
	require.Equal(t, pbUG0, resAll.Ugs[0])

	resUpdate, err := client.UpdateUserGroup(ctx, &pb.UpdateUGRequest{
		Ug: &pb.UserGroup{Id: pbUG0.Id, Name: "user_group0"},
	})
	require.NoError(t, err)
	newPbUG0 := resUpdate.Ug
	require.Equal(t, pbUG0.Id, newPbUG0.Id)
	require.Equal(t, "user_group0", newPbUG0.Name)

	resAll, err = client.UserGroups(ctx, &pb.AllUGsRequest{})
	require.NoError(t, err)
	require.Len(t, resAll.Ugs, 1)
	require.Equal(t, newPbUG0, resAll.Ugs[0])

	pbUG0 = newPbUG0

	_, err = client.DeleteUserGroup(ctx, &pb.DeleteUGRequest{Id: pbUG0.Id})
	require.NoError(t, err)

	resAll, err = client.UserGroups(ctx, &pb.AllUGsRequest{})
	require.NoError(t, err)
	require.Len(t, resAll.Ugs, 0)
}

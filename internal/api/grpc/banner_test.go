package grpc //nolint:dupl

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
)

func TestCRUDBanner(t *testing.T) {
	apiHost := os.Getenv("ADR_TEST_API_HOST")
	if apiHost == "" {
		t.Skipf("Skip integration test")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, apiHost, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	client := pb.NewBannersClient(conn)
	resAll, err := client.Banners(ctx, &pb.AllBannersRequest{})
	require.NoError(t, err)
	require.Empty(t, resAll.Banners)

	resCreate, err := client.CreateBanner(ctx, &pb.CreateBannerRequest{
		Banner: &pb.CreateBannerRequest_NewBanner{Name: "baaner0"},
	})
	require.NoError(t, err)
	pbBanner0 := resCreate.Banner
	require.Greater(t, int(pbBanner0.Id), 0)
	require.Equal(t, "baaner0", pbBanner0.Name)

	resAll, err = client.Banners(ctx, &pb.AllBannersRequest{})
	require.NotNil(t, resAll)
	require.NoError(t, err)
	require.Len(t, resAll.Banners, 1)
	require.Equal(t, pbBanner0, resAll.Banners[0])

	resUpdate, err := client.UpdateBanner(ctx, &pb.UpdateBannerRequest{
		Banner: &pb.Banner{Id: pbBanner0.Id, Name: "banner0"},
	})
	require.NoError(t, err)
	newPbBanner0 := resUpdate.Banner
	require.Equal(t, pbBanner0.Id, newPbBanner0.Id)
	require.Equal(t, "banner0", newPbBanner0.Name)

	resAll, err = client.Banners(ctx, &pb.AllBannersRequest{})
	require.NoError(t, err)
	require.Len(t, resAll.Banners, 1)
	require.Equal(t, newPbBanner0, resAll.Banners[0])

	pbBanner0 = newPbBanner0

	_, err = client.DeleteBanner(ctx, &pb.DeleteBannerRequest{Id: pbBanner0.Id})
	require.NoError(t, err)

	resAll, err = client.Banners(ctx, &pb.AllBannersRequest{})
	require.NoError(t, err)
	require.Len(t, resAll.Banners, 0)
}

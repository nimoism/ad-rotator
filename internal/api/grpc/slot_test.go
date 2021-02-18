package grpc //nolint:dupl

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
)

func TestCRUDSlot(t *testing.T) {
	apiHost := os.Getenv("ADR_TEST_API_HOST")
	if apiHost == "" {
		t.Skipf("Skip integration test")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, apiHost, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	client := pb.NewSlotsClient(conn)
	resAll, err := client.Slots(ctx, &pb.AllSlotsRequest{})
	require.NoError(t, err)
	require.Empty(t, resAll.Slots)

	resCreate, err := client.CreateSlot(ctx, &pb.CreateSlotRequest{
		Slot: &pb.CreateSlotRequest_NewSlot{Name: "solt0"},
	})
	require.NoError(t, err)
	pbSlot0 := resCreate.Slot
	require.Greater(t, int(pbSlot0.Id), 0)
	require.Equal(t, "solt0", pbSlot0.Name)

	resAll, err = client.Slots(ctx, &pb.AllSlotsRequest{})
	require.NotNil(t, resAll)
	require.NoError(t, err)
	require.Len(t, resAll.Slots, 1)
	require.Equal(t, pbSlot0, resAll.Slots[0])

	resUpdate, err := client.UpdateSlot(ctx, &pb.UpdateSlotRequest{
		Slot: &pb.Slot{Id: pbSlot0.Id, Name: "slot0"},
	})
	require.NoError(t, err)
	newPbSlot0 := resUpdate.Slot
	require.Equal(t, pbSlot0.Id, newPbSlot0.Id)
	require.Equal(t, "slot0", newPbSlot0.Name)

	resAll, err = client.Slots(ctx, &pb.AllSlotsRequest{})
	require.NoError(t, err)
	require.Len(t, resAll.Slots, 1)
	require.Equal(t, newPbSlot0, resAll.Slots[0])

	pbSlot0 = newPbSlot0

	_, err = client.DeleteSlot(ctx, &pb.DeleteSlotRequest{Id: pbSlot0.Id})
	require.NoError(t, err)

	resAll, err = client.Slots(ctx, &pb.AllSlotsRequest{})
	require.NoError(t, err)
	require.Len(t, resAll.Slots, 0)
}

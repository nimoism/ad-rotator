package grpc

import (
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/nimoism/ad-rotator/internal/api/grpc/pb"
)

func initTestLayout(ctx context.Context, conn grpc.ClientConnInterface) error {
	bannersClient := pb.NewBannersClient(conn)
	slotsClient := pb.NewSlotsClient(conn)
	ugClient := pb.NewUserGroupsClient(conn)

	bannersNames := []string{"banner0", "banner1", "banner2"}
	pbBanners := make([]*pb.Banner, 0, len(bannersNames))

	for _, name := range bannersNames {
		res, err := bannersClient.CreateBanner(ctx, &pb.CreateBannerRequest{
			Banner: &pb.CreateBannerRequest_NewBanner{Name: name},
		})
		if err != nil {
			return err
		}
		pbBanners = append(pbBanners, res.Banner)
	}

	slotsNames := []string{"slot0", "slot1"}
	pbSlots := make([]*pb.Slot, 0, len(slotsNames))
	for _, name := range slotsNames {
		res, err := slotsClient.CreateSlot(ctx, &pb.CreateSlotRequest{
			Slot: &pb.CreateSlotRequest_NewSlot{Name: name},
		})
		if err != nil {
			return err
		}
		pbSlots = append(pbSlots, res.Slot)
	}

	ugsNames := []string{"0-18", "18-35", "35+"}
	for _, name := range ugsNames {
		_, err := ugClient.CreateUserGroup(ctx, &pb.CreateUGRequest{
			Ug: &pb.CreateUGRequest_NewUserGroup{Name: name},
		})
		if err != nil {
			return err
		}
	}

	for _, bs := range []struct {
		sI int
		bI int
	}{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
	} {
		_, err := bannersClient.BindSlot(ctx, &pb.BindSlotRequest{
			SlotId:   pbSlots[bs.sI].Id,
			BannerId: pbBanners[bs.bI].Id,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func destroyTestLayout(ctx context.Context, conn grpc.ClientConnInterface) error {
	bannersClient := pb.NewBannersClient(conn)
	slotsClient := pb.NewSlotsClient(conn)
	ugClient := pb.NewUserGroupsClient(conn)

	resBanners, err := bannersClient.Banners(ctx, &pb.AllBannersRequest{})
	if err != nil {
		return err
	}
	for _, banner := range resBanners.Banners {
		if _, err = bannersClient.DeleteBanner(ctx, &pb.DeleteBannerRequest{Id: banner.Id}); err != nil {
			return err
		}
	}

	resSlots, err := slotsClient.Slots(ctx, &pb.AllSlotsRequest{})
	if err != nil {
		return err
	}
	for _, slot := range resSlots.Slots {
		if _, err := slotsClient.DeleteSlot(ctx, &pb.DeleteSlotRequest{Id: slot.Id}); err != nil {
			return err
		}
	}

	resUGs, err := ugClient.UserGroups(ctx, &pb.AllUGsRequest{})
	if err != nil {
		return err
	}
	for _, ug := range resUGs.Ugs {
		if _, err := ugClient.DeleteUserGroup(ctx, &pb.DeleteUGRequest{Id: ug.Id}); err != nil {
			return err
		}
	}
	return nil
}

func TestBannersShows(t *testing.T) {
	apiHost := os.Getenv("ADR_TEST_API_HOST")
	if apiHost == "" {
		t.Skipf("Skip integration test")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, apiHost, grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()

	err = initTestLayout(ctx, conn)
	require.NoError(t, err)
	defer func() {
		if err := destroyTestLayout(ctx, conn); err != nil {
			t.Error(err)
		}
	}()

	bClient := pb.NewBannersClient(conn)
	sClient := pb.NewSlotsClient(conn)
	ugClient := pb.NewUserGroupsClient(conn)

	resBanners, err := bClient.Banners(ctx, &pb.AllBannersRequest{})
	require.NoError(t, err)
	banners := resBanners.Banners

	resSlots, err := sClient.Slots(ctx, &pb.AllSlotsRequest{})
	require.NoError(t, err)
	slots := resSlots.Slots

	resUGs, err := ugClient.UserGroups(ctx, &pb.AllUGsRequest{})
	require.NoError(t, err)
	ugs := resUGs.Ugs

	bannersIDsIdxs := make(map[int64]int, len(banners))
	for i, banner := range banners {
		bannersIDsIdxs[banner.Id] = i
	}
	require.Len(t, banners, 3)
	shows := make([]int, len(banners))
	lucks := []float32{0.2, 0.8, 0.2}
	topIdx := 1
	slotID := slots[0].Id
	ugID := ugs[0].Id

	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)
	rand.Seed(seed)
	for i := 0; i < 1000; i++ {
		res, err := bClient.Banner(ctx, &pb.BannerRequest{SlotId: slotID, UserGroupId: ugID})
		require.NoError(t, err)
		banner := res.Banner
		bannerIdx := bannersIDsIdxs[banner.Id]
		shows[bannerIdx]++
		if rand.Float32() < lucks[bannerIdx] { //nolint:gosec
			_, err = bClient.Click(ctx, &pb.ClickRequest{BannerId: banner.Id, SlotId: slotID, UserGroupId: ugID})
			require.NoError(t, err)
		}
	}

	var maxIdx, maxCount int
	for i, count := range shows {
		if maxCount < count {
			maxCount, maxIdx = count, i
		}
	}
	require.Equal(t, topIdx, maxIdx)
}

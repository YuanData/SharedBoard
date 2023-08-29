package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/YuanData/SharedBoard/db/mock"
	db "github.com/YuanData/SharedBoard/db/sqlc"
	"github.com/YuanData/SharedBoard/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomSharedlink() db.Sharedlink {
	return db.Sharedlink{
		ID:      util.RandomInt(1, 1000),
		Name:    util.RandomSentence(),
		Urlhash: util.RandomUUID(),
	}
}
func TestGetSharedlinkAPI(t *testing.T) {
	sharedlink := randomSharedlink()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().
		GetSharedlink(gomock.Any(), gomock.Eq(sharedlink.ID)).
		Times(1).
		Return(sharedlink, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/sharedlink/%d", sharedlink.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchSharedlink(t, recorder.Body, sharedlink)
}

func requireBodyMatchSharedlink(t *testing.T, body *bytes.Buffer, trader db.Sharedlink) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var gotSharedlink db.Sharedlink
	err = json.Unmarshal(data, &gotSharedlink)
	require.NoError(t, err)
	require.Equal(t, trader, gotSharedlink)
}

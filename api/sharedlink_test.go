package api

import (
	"bytes"
	"database/sql"
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
	testCases := []struct {
		name          string
		SharedlinkID  int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:         "OK",
			SharedlinkID: sharedlink.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSharedlink(gomock.Any(), gomock.Eq(sharedlink.ID)).
					Times(1).
					Return(sharedlink, nil)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchSharedlink(t, recorder.Body, sharedlink)
			},
		},
		{
			name:         "NotFound",
			SharedlinkID: sharedlink.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSharedlink(gomock.Any(), gomock.Eq(sharedlink.ID)).
					Times(1).
					Return(db.Sharedlink{}, sql.ErrNoRows)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:         "InternalError",
			SharedlinkID: sharedlink.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSharedlink(gomock.Any(), gomock.Eq(sharedlink.ID)).
					Times(1).
					Return(db.Sharedlink{}, sql.ErrConnDone)

			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:         "InvalidID",
			SharedlinkID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSharedlink(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/sharedlink/%d", tc.SharedlinkID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchSharedlink(t *testing.T, body *bytes.Buffer, trader db.Sharedlink) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)
	var gotSharedlink db.Sharedlink
	err = json.Unmarshal(data, &gotSharedlink)
	require.NoError(t, err)
	require.Equal(t, trader, gotSharedlink)
}

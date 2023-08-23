package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	recaptchaToken = "03ADUVZwA_txr-7UFRUYM6VcSWHL-8ADyHRwzQKUEPIqJKTU_a71lUuOKi8asQaFz0Sz8xPeGA3h2EX-ACD7btUq7Jt6XH4lKwCExGo2uGF3CyvhuI6_JvphIG00yQkwmadksTz4Hl7YX4V_-B9dD60rK-uQZ7q0JLwzev97mN5OQBp3A5rYpMGp2RlFWFvUiSZazgunK4nhsJKgpjp-uVB89mFAzQKGHAyE4TqmaqFZ0ntfFhRBEaEKTpXc7gMX-5glnqkf86l0ufQqhEpehrXmXwa5GgPaGKRe-5CHrV4uGeV1VG_sb2RaeskPUgGCNf20ykpkGOE3y-sLEES9mXOJ5M6MD5jsms6LaBDX7cOPcRcaWfEAyogEhS9m0dOe46VY6YmiAqQCWaOCYxBFhdeM8e50OXKex8iUJrwckB5Apwg3jt3tICCb_VYEPDaYolsYs5HF0AAANH1cvnp4Orjonfv0TScopB_9d2WqIOeJvOhd9XSw7zUgm_x49cZ9MlY3FTU-7R-PnuehumiPwPeNC8CmqAuM46avve5XqQhUDOQF4qpb0p5RT7tj-PYSVAl4N6KMzWjTpTAhIXpRgqcAs2ImGMmvl4gH0I3t_N0DlobUpEFrY8NEkj7wdvMz1Ezed4MfjXlhYzfTehYkNAcPF6DyKXhV_Jv0WohZcdZmn_W-7izOlMY3l_mVDaDralQ53PU6eUqkG8pWgBa_tckcIuzJCTndAYnAAnkcqu5QR9LmC7299H2itRNftKvN8kmFX7SqmX6ZNDoJETiHu4AA7YSb8DrYGyeZft7RivyTsjwEnrrq5UdKbzpV8PEO3pwDyQ8SAGU8kgMyhS3rc-YuVIeVRJE3A9CFGf9J2gvdZ6ZXnoAYKxShK-0smvq2zLHvdsK_U2hBndPM_FHc-NWSGfNwB_qkZIyzdPO9gDEqNVgP-sbt3ozOQpNE4_5eakf1tvUWOU1DkaerxQRMdiUpGH4oEfB2070NULitxNcdF4xVVkhUQ8-uO9k5Xy5G8clnGL0SgRLQ3IucD--TtGFBsjNhtQMywvoCPaJTUpXg343Ky_qvj3HwnHfyMYM8VKaRpQs3EKdYAlmLJbDVqfShZnPvkHMwE_rqDVDeX2phVaH-SU2Yz2Zf8pP64WMA7SD5kC0QbOMQpMvi-UMxQ5WsJ0cyY4tzm3B0K3immxJtAeeZlSzekMxj-EXFBvpdh7KBhmlivtJJMxAyVW0KGcYsZvO5dqRMcpes9b7nUV1n_m2rQfq7tbjdEWOV6ELCmfkrKmf3U0vON6yN7Y_bVU_GNOkAERx-mDqVEcOGVe0nRfq0YhCir_vfX194liawKXEj1BYRtpyi0vCMyr4MhDV1qY89gRWXMQ59Hj3QjydC9gMnihSnqW-xFeVMs3KCcmvpVsj-VJERSRNawgWKBAmdZAxU0CIXKzV-gV-COwUmrGOIxBHnDTE4Hij3nJzNGZ1raG5gTyWah72vvMyUlqiGALf1zGCU2FPTKW_rgdAkCQdI4TqacWpMn-cobT7AQt0XOzT8XOcqhtQLQG40aiFbhXZygXLp98yWb2_Kftj5aZozA0LGmLoh6TeUP3rmSh1gegdNf_VY7SocIqEAw4-YX8QMvWh-lYY6EXtggpfbHkshEQtYnG9kdQIY6LECPJIVjLSoYim4OVcfD_yTayQR15-XiSOe-rpyAVDPhQlWFiQxbyc7wm-H1BKCpyGRpnTbj9IZFKsUVk7u_x1X4Fls13-OOJXrcq59JvW-La8BUPjhuwte-Gc04"
	cont           = "https://open.spotify.com/__noul__?l2l=1&nd=1&flow_ctx=bb7699a3-77a9-47ea-98f3-5c4e293c602d%3A1692816742"
	flowCtx        = "bb7699a3-77a9-47ea-98f3-5c4e293c602d:1692816742"
)

type SpotifyResp struct {
	Href  string `json:"href"`
	Items []Item `json:"items"`
}
type Item struct {
	Track SpotifySong `json:"track"`
}
type SpotifySong struct {
	Id      string   `json:"id"`
	Artists []Artist `json:"artists"`
	Name    string   `json:"name"`
}
type Artist struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func getJwt(user User) (JWT string, err error) {
	url := "https://accounts.spotify.com/login/password"
	data := struct {
		User
		Remember bool   `json:"remember"`
		Continue string `json:"continue"`
		FlowCtx  string `json:"flowCtx"`
	}{
		User{
			user.username,
			user.password,
		},
		true,
		cont,
		flowCtx,
	}
	dataBytes, err := json.Marshal(&data)
	if err != nil {
		return "", err

	}
	reader := bytes.NewReader(dataBytes)
	res, err := http.Post(url, "application/json", reader)
	// Close response body
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	responseBody, err := io.ReadAll(res.Body)
	fmt.Println(string(responseBody))
	return "", nil
}
func fetchSongs(jwt, album string) (*[]SpotifySong, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", album)
	fmt.Println(url)
	var spotifyResp SpotifyResp
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", jwt)
	client := http.Client{}
	res, err := client.Do(req)
	fmt.Println(jwt)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&spotifyResp)
	json.Unmarshal(b, &spotifyResp)
	fmt.Println(string(b))
	fmt.Println(spotifyResp)
	return nil, nil
}

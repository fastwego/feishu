
func Upload(ctx *feishu.App, content string, params url.Values) (resp []byte, err error) {

	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("content", path.Base(content))
		if err != nil {
			return
		}
		file, err := os.Open(content)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

		// field
		err = m.WriteField("type", params.Get("type"))
		if err != nil {
			return
		}

		err = m.WriteField("name", path.Base(content))
		if err != nil {
			return
		}

	}()

	accessToken, err := ctx.GetTenantAccessTokenHandler()
	if err != nil {
		return
	}
	header := http.Header{}
	header.Set("Authorization", "Bearer "+accessToken)
	header.Set("Content-Type", m.FormDataContentType())

	return ctx.Client.HTTPPost(feishu.FeishuServerUrl2+apiUpload, r, header)
}
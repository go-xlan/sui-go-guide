package suiapi

type TxBytesMessage struct {
	TxBytes string `json:"txBytes"`
}

type DigestMessage struct {
	Digest string `json:"digest"`
}

type EffectsStatusStatusMessage struct {
	Effects struct {
		Status struct {
			Status string `json:"status"`
		} `json:"status"`
	} `json:"effects"`
}

type CoinType struct {
	Balance             string `json:"balance"`
	CoinObjectId        string `json:"coinObjectId"`
	CoinType            string `json:"coinType"`
	Digest              string `json:"digest"`
	PreviousTransaction string `json:"previousTransaction"`
	Version             string `json:"version"`
}

type CoinMetadata struct {
	Decimals    int    `json:"decimals"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
}

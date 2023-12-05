package order

/**
 * @Author: bitget-sdk-team
 * @Date: 2022-09-30 10:46
 * @DES: cancel the order request
 */
type CancelOrderReq struct {
	/**
	 * Currency pair
	 */
	Symbol string `json:"symbol"`
	/**
	 * Order Id
	 */
	OrderId string `json:"orderId"`
}

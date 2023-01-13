/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-26 10:33:53
 * @LastEditTime: 2023-01-13 14:43:53
 * @Description: Do not edit
 */
package query

// Herte handler query binding struct
type PaginationQuery struct {
	Offset uint   `form:"offset" json:"offset" query:"offset"` //偏移量[页数] example(1)
	Limit  uint   `form:"limit" json:"limit" query:"limit"`    //单页限制 example(10)
	Stime  string `json:"stime" form:"stime" query:"stime"`    //开始时间 example("2016-01-02 03:04:05")
	Etime  string `json:"etime" form:"etime" query:"etime"`    //结束时间 example("2016-01-02 03:04:05")
}

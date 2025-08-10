package main

import (
	"fmt"
	"time"
)

type PaymentStratey interface {
	pay(amount float64) string
	refund(amount float64) string
}

type CreditCardPayment struct {
	cardId string
}

func newCreditCardPayment(cardId string) *CreditCardPayment {
	return &CreditCardPayment{
		cardId: cardId,
	}
}

func (c *CreditCardPayment) pay(amount float64) string {
	return fmt.Sprintf("")
}

func (c *CreditCardPayment) refund(amount float64) string {
	return fmt.Sprintf("")
}

type OtherPayment struct {
	pay_name string
}

/*
	实现2个接口
*/

func paymentFactory(payment_type string, details ...string) PaymentStratey {
	switch payment_type {
	case "credit":
		return newCreditCardPayment(details[0])
	case "other":
		return nil //
	default:
		// panic是一种用于处理不可恢复错误的机制。当程序遇到严重问题时，
		// 可以通过panic中断当前流程并输出错误信息，同时触发调用栈的回
		panic("异常类型")
		// return nil
	}
}

type User struct {
	user_name string
	payment   PaymentStratey
	orders    []*Order
}

type Order struct {
	order_num   string
	user        *User
	product     []*Product
	num         int
	total_price float64
	status      string
	create_time time.Time
}

type Product struct {
	p_name    string
	price     float64
	quantity  int // 单笔数量
	inventory int // 库存
}

func newUser(name string) *User {
	return &User{
		user_name: name,
	}
}

func (u *User) setPayment(ps PaymentStratey) {
	u.payment = ps
}

func (u *User) placeOrder(products []*Product) *Order {
	order := &Order{
		order_num:   "",
		user:        u,
		product:     products,
		status:      "created",
		create_time: time.Now(),
	}

	/* 计算商品价格  锁定库存 */
	u.orders = append(u.orders, order)
	return order
}

//  struct 相当于一个类
// func (c *CreditCardPayment) refund (amount float64) string { 相当于 给这个类 加个 方法
// type PaymentStratey interface {  定义一个接口
// go 中 继承 是组合的方式实现
// 没有抽象类
// 各个模块可以灵活 嫁接 使用不同的拼装方式 实现起来灵活多样

// Code generated by gowsdl DO NOT EDIT.

package sms

import (
	"context"
	"encoding/xml"
	"time"

	"mailer/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type GetMessageState struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetMessageState"`

	SessionID string `xml:"sessionID,omitempty" json:"sessionID,omitempty"`

	MessageID int64 `xml:"messageID,omitempty" json:"messageID,omitempty"`
}

type GetMessageStateResponse struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetMessageStateResponse"`

	GetMessageStateResult *SmsMessageStateInfo `xml:"GetMessageStateResult,omitempty" json:"GetMessageStateResult,omitempty"`
}

type GetSessionID struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetSessionID"`

	Login string `xml:"login,omitempty" json:"login,omitempty"`

	Password string `xml:"password,omitempty" json:"password,omitempty"`
}

type GetSessionIDResponse struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetSessionIDResponse"`

	GetSessionIDResult string `xml:"GetSessionIDResult,omitempty" json:"GetSessionIDResult,omitempty"`
}

type SendMessage struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com SendMessage"`

	SessionID string `xml:"sessionID,omitempty" json:"sessionID,omitempty"`

	Message *Message `xml:"message,omitempty" json:"message,omitempty"`
}

type SendMessageResponse struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com SendMessageResponse"`

	SendMessageResult *ArrayOfString `xml:"SendMessageResult,omitempty" json:"SendMessageResult,omitempty"`
}

type SendMessageByTimeZone struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com SendMessageByTimeZone"`

	SessionID string `xml:"sessionID,omitempty" json:"sessionID,omitempty"`

	SourceAddress string `xml:"sourceAddress,omitempty" json:"sourceAddress,omitempty"`

	DestinationAddress string `xml:"destinationAddress,omitempty" json:"destinationAddress,omitempty"`

	Data string `xml:"data,omitempty" json:"data,omitempty"`

	SendDate time.Time `xml:"sendDate,omitempty" json:"sendDate,omitempty"`

	Validity int32 `xml:"validity,omitempty" json:"validity,omitempty"`
}

type SendMessageByTimeZoneResponse struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com SendMessageByTimeZoneResponse"`

	SendMessageByTimeZoneResult *ArrayOfString `xml:"SendMessageByTimeZoneResult,omitempty" json:"SendMessageByTimeZoneResult,omitempty"`
}

type SendMessageByTimeZoneToAddresses struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com SendMessageByTimeZoneToAddresses"`

	SessionID string `xml:"sessionID,omitempty" json:"sessionID,omitempty"`

	SourceAddress string `xml:"sourceAddress,omitempty" json:"sourceAddress,omitempty"`

	DestinationAddresses *ArrayOfString `xml:"destinationAddresses,omitempty" json:"destinationAddresses,omitempty"`

	Data string `xml:"data,omitempty" json:"data,omitempty"`

	SendDate time.Time `xml:"sendDate,omitempty" json:"sendDate,omitempty"`

	Validity int32 `xml:"validity,omitempty" json:"validity,omitempty"`
}

type SendMessageByTimeZoneToAddressesResponse struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com SendMessageByTimeZoneToAddressesResponse"`

	SendMessageByTimeZoneToAddressesResult *ArrayOfString `xml:"SendMessageByTimeZoneToAddressesResult,omitempty" json:"SendMessageByTimeZoneToAddressesResult,omitempty"`
}

type GetBalance struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetBalance"`

	SessionID string `xml:"sessionID,omitempty" json:"sessionID,omitempty"`
}

type GetBalanceResponse struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetBalanceResponse"`

	GetBalanceResult float64 `xml:"GetBalanceResult,omitempty" json:"GetBalanceResult,omitempty"`
}

type GetIncomingMessages struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetIncomingMessages"`

	SessionID string `xml:"sessionID,omitempty" json:"sessionID,omitempty"`

	MinDateUTC time.Time `xml:"minDateUTC,omitempty" json:"minDateUTC,omitempty"`

	MaxDateUTC time.Time `xml:"maxDateUTC,omitempty" json:"maxDateUTC,omitempty"`
}

type GetIncomingMessagesResponse struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetIncomingMessagesResponse"`

	GetIncomingMessagesResult *ArrayOfIncomingMessage `xml:"GetIncomingMessagesResult,omitempty" json:"GetIncomingMessagesResult,omitempty"`
}

type GetStatistics struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetStatistics"`

	SessionID string `xml:"sessionID,omitempty" json:"sessionID,omitempty"`

	StartDateTime time.Time `xml:"startDateTime,omitempty" json:"startDateTime,omitempty"`

	EndDateTime time.Time `xml:"endDateTime,omitempty" json:"endDateTime,omitempty"`
}

type GetStatisticsResponse struct {
	XMLName xml.Name `xml:"http://ws.devinosms.com GetStatisticsResponse"`

	GetStatisticsResult *DeliveryStatistics `xml:"GetStatisticsResult,omitempty" json:"GetStatisticsResult,omitempty"`
}

type SmsMessageStateInfo struct {
	CreationDateUtc string `xml:"CreationDateUtc,omitempty" json:"CreationDateUtc,omitempty"`

	SubmittedDateUtc string `xml:"SubmittedDateUtc,omitempty" json:"SubmittedDateUtc,omitempty"`

	ReportedDateUtc string `xml:"ReportedDateUtc,omitempty" json:"ReportedDateUtc,omitempty"`

	TimeStampUtc string `xml:"TimeStampUtc,omitempty" json:"TimeStampUtc,omitempty"`

	State int32 `xml:"State,omitempty" json:"State,omitempty"`

	StateDescription string `xml:"StateDescription,omitempty" json:"StateDescription,omitempty"`

	Price float64 `xml:"Price,omitempty" json:"Price,omitempty"`
}

type Message struct {
	Data string `xml:"Data,omitempty" json:"Data,omitempty"`

	DelayUntilUtc time.Time `xml:"DelayUntilUtc,omitempty" json:"DelayUntilUtc,omitempty"`

	DestinationAddresses *ArrayOfString `xml:"DestinationAddresses,omitempty" json:"DestinationAddresses,omitempty"`

	SourceAddress string `xml:"SourceAddress,omitempty" json:"SourceAddress,omitempty"`

	ReceiptRequested bool `xml:"ReceiptRequested,omitempty" json:"ReceiptRequested,omitempty"`

	Validity int32 `xml:"Validity,omitempty" json:"Validity,omitempty"`

	Optionals *ArrayOfString `xml:"Optionals,omitempty" json:"Optionals,omitempty"`
}

type ArrayOfString struct {
	Astring []string `xml:"string,omitempty" json:"string,omitempty"`
}

type ArrayOfIncomingMessage struct {
	IncomingMessage []*IncomingMessage `xml:"IncomingMessage,omitempty" json:"IncomingMessage,omitempty"`
}

type IncomingMessage struct {
	ID int64 `xml:"ID,omitempty" json:"ID,omitempty"`

	Data string `xml:"Data,omitempty" json:"Data,omitempty"`

	SourceAddress string `xml:"SourceAddress,omitempty" json:"SourceAddress,omitempty"`

	DestinationAddress string `xml:"DestinationAddress,omitempty" json:"DestinationAddress,omitempty"`

	CreatedDateUtc time.Time `xml:"CreatedDateUtc,omitempty" json:"CreatedDateUtc,omitempty"`
}

type DeliveryStatistics struct {
	Sent int32 `xml:"Sent,omitempty" json:"Sent,omitempty"`

	Delivered int32 `xml:"Delivered,omitempty" json:"Delivered,omitempty"`

	Errors int32 `xml:"Errors,omitempty" json:"Errors,omitempty"`

	InProcess int32 `xml:"InProcess,omitempty" json:"InProcess,omitempty"`

	Expired int32 `xml:"Expired,omitempty" json:"Expired,omitempty"`

	Rejected int32 `xml:"Rejected,omitempty" json:"Rejected,omitempty"`

	Optimized int32 `xml:"Optimized,omitempty" json:"Optimized,omitempty"`

	Tarificated int32 `xml:"Tarificated,omitempty" json:"Tarificated,omitempty"`

	Total int32 `xml:"Total,omitempty" json:"Total,omitempty"`

	TotalWithErrors int32 `xml:"TotalWithErrors,omitempty" json:"TotalWithErrors,omitempty"`

	DeliveryRatio float64 `xml:"DeliveryRatio,omitempty" json:"DeliveryRatio,omitempty"`
}

type SmsServiceSoap interface {

	/* Возвращает статус сообщения и время обновления статуса */
	GetMessageState(request *GetMessageState) (*GetMessageStateResponse, error)

	GetMessageStateContext(ctx context.Context, request *GetMessageState) (*GetMessageStateResponse, error)

	/* Возвращает идентификатор сессии для пользователя */
	GetSessionID(request *GetSessionID) (*GetSessionIDResponse, error)

	GetSessionIDContext(ctx context.Context, request *GetSessionID) (*GetSessionIDResponse, error)

	/* Отправляет сообщение адресатам и возвращает системные идентификаторы сообщений */
	SendMessage(request *SendMessage) (*SendMessageResponse, error)

	SendMessageContext(ctx context.Context, request *SendMessage) (*SendMessageResponse, error)

	/* Отправляет сообщение учитывая часовой пояс получателя */
	SendMessageByTimeZone(request *SendMessageByTimeZone) (*SendMessageByTimeZoneResponse, error)

	SendMessageByTimeZoneContext(ctx context.Context, request *SendMessageByTimeZone) (*SendMessageByTimeZoneResponse, error)

	/* Отправляет сообщение учитывая часовой пояс получателя на указанные адреса получателей */
	SendMessageByTimeZoneToAddresses(request *SendMessageByTimeZoneToAddresses) (*SendMessageByTimeZoneToAddressesResponse, error)

	SendMessageByTimeZoneToAddressesContext(ctx context.Context, request *SendMessageByTimeZoneToAddresses) (*SendMessageByTimeZoneToAddressesResponse, error)

	/* Возвращает баланс пользователя текущей сессии */
	GetBalance(request *GetBalance) (*GetBalanceResponse, error)

	GetBalanceContext(ctx context.Context, request *GetBalance) (*GetBalanceResponse, error)

	/* Возвращает входящие сообщения пользователя текущей сессии */
	GetIncomingMessages(request *GetIncomingMessages) (*GetIncomingMessagesResponse, error)

	GetIncomingMessagesContext(ctx context.Context, request *GetIncomingMessages) (*GetIncomingMessagesResponse, error)

	/* Возвращает статистику отправленных смс за заданный промежуток времени */
	GetStatistics(request *GetStatistics) (*GetStatisticsResponse, error)

	GetStatisticsContext(ctx context.Context, request *GetStatistics) (*GetStatisticsResponse, error)
}

type smsServiceSoap struct {
	client *soap.Client
}

func NewSmsServiceSoap(client *soap.Client) SmsServiceSoap {
	return &smsServiceSoap{
		client: client,
	}
}

func (service *smsServiceSoap) GetMessageStateContext(ctx context.Context, request *GetMessageState) (*GetMessageStateResponse, error) {
	response := new(GetMessageStateResponse)
	err := service.client.CallContext(ctx, "http://ws.devinosms.com/GetMessageState", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *smsServiceSoap) GetMessageState(request *GetMessageState) (*GetMessageStateResponse, error) {
	return service.GetMessageStateContext(
		context.Background(),
		request,
	)
}

func (service *smsServiceSoap) GetSessionIDContext(ctx context.Context, request *GetSessionID) (*GetSessionIDResponse, error) {
	response := new(GetSessionIDResponse)
	err := service.client.CallContext(ctx, "http://ws.devinosms.com/GetSessionID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *smsServiceSoap) GetSessionID(request *GetSessionID) (*GetSessionIDResponse, error) {
	return service.GetSessionIDContext(
		context.Background(),
		request,
	)
}

func (service *smsServiceSoap) SendMessageContext(ctx context.Context, request *SendMessage) (*SendMessageResponse, error) {
	response := new(SendMessageResponse)
	err := service.client.CallContext(ctx, "http://ws.devinosms.com/SendMessage", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *smsServiceSoap) SendMessage(request *SendMessage) (*SendMessageResponse, error) {
	return service.SendMessageContext(
		context.Background(),
		request,
	)
}

func (service *smsServiceSoap) SendMessageByTimeZoneContext(ctx context.Context, request *SendMessageByTimeZone) (*SendMessageByTimeZoneResponse, error) {
	response := new(SendMessageByTimeZoneResponse)
	err := service.client.CallContext(ctx, "http://ws.devinosms.com/SendMessageByTimeZone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *smsServiceSoap) SendMessageByTimeZone(request *SendMessageByTimeZone) (*SendMessageByTimeZoneResponse, error) {
	return service.SendMessageByTimeZoneContext(
		context.Background(),
		request,
	)
}

func (service *smsServiceSoap) SendMessageByTimeZoneToAddressesContext(ctx context.Context, request *SendMessageByTimeZoneToAddresses) (*SendMessageByTimeZoneToAddressesResponse, error) {
	response := new(SendMessageByTimeZoneToAddressesResponse)
	err := service.client.CallContext(ctx, "http://ws.devinosms.com/SendMessageByTimeZoneToAddresses", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *smsServiceSoap) SendMessageByTimeZoneToAddresses(request *SendMessageByTimeZoneToAddresses) (*SendMessageByTimeZoneToAddressesResponse, error) {
	return service.SendMessageByTimeZoneToAddressesContext(
		context.Background(),
		request,
	)
}

func (service *smsServiceSoap) GetBalanceContext(ctx context.Context, request *GetBalance) (*GetBalanceResponse, error) {
	response := new(GetBalanceResponse)
	err := service.client.CallContext(ctx, "http://ws.devinosms.com/GetBalance", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *smsServiceSoap) GetBalance(request *GetBalance) (*GetBalanceResponse, error) {
	return service.GetBalanceContext(
		context.Background(),
		request,
	)
}

func (service *smsServiceSoap) GetIncomingMessagesContext(ctx context.Context, request *GetIncomingMessages) (*GetIncomingMessagesResponse, error) {
	response := new(GetIncomingMessagesResponse)
	err := service.client.CallContext(ctx, "http://ws.devinosms.com/GetIncomingMessages", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *smsServiceSoap) GetIncomingMessages(request *GetIncomingMessages) (*GetIncomingMessagesResponse, error) {
	return service.GetIncomingMessagesContext(
		context.Background(),
		request,
	)
}

func (service *smsServiceSoap) GetStatisticsContext(ctx context.Context, request *GetStatistics) (*GetStatisticsResponse, error) {
	response := new(GetStatisticsResponse)
	err := service.client.CallContext(ctx, "http://ws.devinosms.com/GetStatistics", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *smsServiceSoap) GetStatistics(request *GetStatistics) (*GetStatisticsResponse, error) {
	return service.GetStatisticsContext(
		context.Background(),
		request,
	)
}

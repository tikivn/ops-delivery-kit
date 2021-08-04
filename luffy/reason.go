package luffy

type Reason string

const NoLongerNeededReason Reason = "no_longer_needed"
const ReceiverReschedule Reason = "receiver_reschedule"
const CanNotContactReason Reason = "can_not_contact"
const WrongAddressReason Reason = "wrong_address"
const DriverRescheduleReason Reason = "driver_reschedule"
const AbsenceOfReceiverReason Reason = "absence_of_receiver"
const OrderAmendmentReason Reason = "order_amendment"
const HandOverReason Reason = "handover"
const LateDeliveryReason Reason = "late_delivery"
const OpenBoxRightLimitedReason Reason = "open_box_right_limited"
const CheckingBoxRightLimited Reason = "checking_box_right_limited"
const WrongProductReason Reason = "wrong_product"
const WrongCodReason Reason = "wrong_cod"
const LowQualityProductReason Reason = "low_quality_product"
const NotPlaceReason Reason = "not_place"
const VirtualCustomerReason Reason = "virtual_customer"
const BoxDamagedReason Reason = "box_damaged"
const OrderDuplicatedReason Reason = "order_duplicated"
const UnavoidableAccidentReason Reason = "unavoidable_accident"
const OutOfZoneReason Reason = "out_of_zone"
const GetOrderAtHubReason Reason = "get_order_at_hub"
const GetParcelAtHub Reason = "get_parcel_at_hub"
const LostParcelReason Reason = "lost_parcel"
const RequestReturnBySenderReason Reason = "request_return_by_sender"
const CanceledReason Reason = "canceled"
const BulkyGoodsReason Reason = "bulky_goods"
const SenderRescheduleReason Reason = "sender_reschedule"
const CustomerNotPrepareCOD Reason = "customer_not_prepare_cod"
const AbsenceOfSender Reason = "absence_of_sender"
const ChangeAddress Reason = "change_address"
const GoodsAreNotReady Reason = "goods_are_not_ready"
const UnsecuredPacking Reason = "unsecured_packing"
const LackOfProductCover Reason = "lack_of_product_cover"
const CanceledBySender Reason = "canceled_by_sender"
const CanNotContactManyTimes Reason = "can_not_contact_many_times"
const AbsenceOfSenderManyTimes Reason = "absence_of_sender__many_times"
const PickedUpByAnother3pl Reason = "picked_up_by_another_3pl"
const SenderOutOfStock Reason = "sender_out_of_stock"
const NeedReturnMoneyBeforePickup Reason = "need_return_money_before_pickup"
const CustomerReturnDirectly Reason = "customer_return_directly"
const GoodIsUsed Reason = "good_is_used"
const BannedGood Reason = "banned_good"
const WrongTelephoneNumber Reason = "wrong_telephone_number"
const SenderBringGoodsToHub Reason = "sender_bring_goods_to_hub"
const DamagedProduct Reason = "damaged_product"
const FailedReturn3Times Reason = "failed_return_3_times"
const WrongContactInfo Reason = "wrong_contact_info"
const WrongInformation Reason = "wrong_information"
const OverWeightGoods Reason = "over_weight_goods"
const FailedPickupManyTimes Reason = "failed_pickup_many_times"
const LateReturn Reason = "late_return"
const CanceledBy3pl Reason = "canceled_by_3pl"
const DriverNotFound Reason = "driver_not_found"
const PackageInReturn Reason = "package_in_return"
const PackageReturned Reason = "package_returned"
const UndefinedReason Reason = "undefined"

func GetReasonName(reason Reason) string {
	switch reason {
	case NoLongerNeededReason:
		return "Khách hàng không còn nhu cầu"
	case ReceiverReschedule:
		return "Khách hàng hẹn giao lại"
	case CanNotContactReason:
		return "Không liên lạc được"
	case WrongAddressReason:
		return "Sai địa chỉ, đổi địa chỉ"
	case DriverRescheduleReason:
		return "Vận chuyển hẹn lại khách hàng"
	case AbsenceOfReceiverReason:
		return "Đến nhà không có người nhận"
	case OrderAmendmentReason:
		return "Yêu cầu sửa đơn hàng"
	case HandOverReason:
		return "Bàn giao cho nhân viên khác"
	case LateDeliveryReason:
		return "Giao hàng chậm quá thời gian cam kết"
	case OpenBoxRightLimitedReason:
		return "Không được kiểm tra hàng"
	case WrongProductReason:
		return "Sai sản phẩm"
	case WrongCodReason:
		return "Sai tiền COD/ Thay đổi thông tin thanh toán"
	case LowQualityProductReason:
		return "Chất lượng sản phẩm kém"
	case NotPlaceReason:
		return "Khách không đặt hàng"
	case VirtualCustomerReason:
		return "Khách hàng ảo"
	case BoxDamagedReason:
		return "Hàng bị hư hỏng"
	case OrderDuplicatedReason:
		return "Đơn hàng bị trùng"
	case UnavoidableAccidentReason:
		return "Trường hợp bất khả kháng"
	case LostParcelReason:
		return "Kiện hàng bị thất lạc"
	case GetOrderAtHubReason:
		return "Khách lấy hàng tại trạm giao"
	case RequestReturnBySenderReason:
		return "Người gửi yêu cầu hoàn hàng"
	case CanceledReason:
		return "Hủy"
	case BulkyGoodsReason:
		return "Hàng cồng kềnh"
	case SenderRescheduleReason:
		return "Khách hẹn"
	default:
		return "Trạng thái không xác định"
	}
}

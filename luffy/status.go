package luffy

// Should move to tms core
type Status string

// Pickup status
const (
	NewStatus              Status = "new"
	ReadyToPickupStatus    Status = "ready_for_picking_up"
	PickingStatus          Status = "picking"
	PickupFailedStatus     Status = "pickup_failed"
	PendingPickupStatus    Status = "pending_pickup"
	CanceledStatus         Status = "canceled"
	SuccessfulPickupStatus Status = "successful_pickup"

	// Delivery status
	ArrivedAtSortingHubStatus         Status = "arrived_at_sorting_hub"
	ArrivedAtTransitHubStatus         Status = "arrived_at_transit_hub"
	ArrivedAtDeliveryHubStatus        Status = "arrived_at_delivery_hub"
	InTransitStatus                   Status = "in_transit"
	StoredAtDeliveryHubStatus         Status = "stored_at_delivery_hub"
	DeliveringStatus                  Status = "delivering"
	DeliveryFailedStatus              Status = "delivery_failed"
	PendingDeliveryStatus             Status = "pending_delivery"
	SuccessfulDeliveryStatus          Status = "successful_delivery"
	ReturnToSenderTriggeredStatus     Status = "return_to_sender_triggered"
	ReturningStatus                   Status = "returning"
	ReturnedStatus                    Status = "returned"
	ContinueDeliveryStatus            Status = "continue_delivery"
	LostStatus                        Status = "lost"
	OnholdStatus                      Status = "onhold"
	ReturnFailedStatus                Status = "return_failed"
	InboundFirstmileHubStatus         Status = "inbound_firstmile_hub"
	OutboundFirstmileHubStatus        Status = "outbound_firstmile_hub"
	InboundSortingHubStatus           Status = "inbound_sorting_hub"
	OutboundSortingHubStatus          Status = "outbound_sorting_hub"
	InboundLastmileHubStatus          Status = "inbound_lastmile_hub"
	InboundDeliveryHubStatus          Status = "inbound_delivery_hub"
	OutboundLastmileHubToReturnStatus Status = "outbound_lastmile_hub_to_return"
	InboundSortingHubToReturnStatus   Status = "inbound_sorting_hub_to_return"
	OutboundSortingHubToReturnStatus  Status = "outbound_sorting_hub_to_return"
	InboundFirstmileHubToReturnStatus Status = "inbound_firstmile_hub_to_return"
	CompletedStatus                   Status = "completed"

	PackedInMasterBox               Status = "packed_in_masterbox"
	ArrivedAtLastMileHub            Status = "arrived_at_lastmile_hub"
	PickupInteractingWithSender     Status = "pickup_interacting_with_sender"
	DeliveryInteractingWithReceiver Status = "delivery_interacting_with_receiver"
	InTransitToReturn               Status = "in_transit_to_return"
	Exception                       Status = "exception"
	Damage                          Status = "damage"

	UndefinedStatus Status = "undefined"
)

func GetStatusName(status Status) string {
	switch status {
	case NewStatus:
		return "Mới"
	case ReadyToPickupStatus:
		return "Sẵn sàng lấy hàng"
	case PickingStatus:
		return "Đang lấy hàng"
	case PickupFailedStatus:
		return "Lấy hàng thất bại"
	case PendingPickupStatus:
		return "Chờ để lấy hàng"
	case ArrivedAtSortingHubStatus:
		return "Tới trạm phân loại hàng"
	case InTransitStatus:
		return "Đang luân chuyển"
	case ArrivedAtTransitHubStatus:
		return "Tới trạm luân chuyển"
	case ArrivedAtDeliveryHubStatus:
		return "Tới trạm giao hàng cho khách"
	case StoredAtDeliveryHubStatus:
		return "Nhập kho vận chuyển"
	case DeliveringStatus:
		return "Trên đường giao hàng"
	case DeliveryFailedStatus:
		return "Giao hàng không thành công"
	case PendingDeliveryStatus:
		return "Chờ để giao hàng"
	case SuccessfulDeliveryStatus:
		return "Giao hàng thành công"
	case SuccessfulPickupStatus:
		return "Lấy hàng thành công"
	case InboundFirstmileHubStatus:
		return "Nhập kho trạm lấy hàng"
	case OutboundFirstmileHubStatus:
		return "Xuất kho trạm lấy hàng"
	case InboundSortingHubStatus:
		return "Nhập kho trạm phân loại hàng"
	case OutboundSortingHubStatus:
		return "Xuất kho trạm phân loại hàng"
	case InboundLastmileHubStatus:
		return "Nhập kho trạm giao hàng"
	case InboundDeliveryHubStatus:
		return "Nhập kho vận chuyển"
	case OutboundLastmileHubToReturnStatus:
		return "Xuất kho trạm giao hàng để trả hàng"
	case InboundSortingHubToReturnStatus:
		return "Nhập kho trạm phân loại để trả hàng"
	case OutboundSortingHubToReturnStatus:
		return "Xuất kho trạm phân loại để trả hàng"
	case InboundFirstmileHubToReturnStatus:
		return "Nhập kho trạm lấy hàng để trả hàng"
	case ReturningStatus:
		return "Trên đường trả hàng"
	case ReturnFailedStatus:
		return "Trả hàng không thành công"
	case ReturnedStatus:
		return "Trả hàng thành công"
	case ContinueDeliveryStatus:
		return "Yêu cầu hoàn hàng"
	case CanceledStatus:
		return "Hủy"
	case LostStatus:
		return "Mất hàng"
	case ReturnToSenderTriggeredStatus:
		return "Yêu cầu hoàn hàng"
	case PackedInMasterBox:
		return "Đã đóng trong Masterbox"
	case ArrivedAtLastMileHub:
		return "Tới trạm giao hàng cho khách"
	case PickupInteractingWithSender:
		return "Tương tác với người gửi"
	case DeliveryInteractingWithReceiver:
		return "Tương tác với người nhận"
	case InTransitToReturn:
		return "Luân chuyển trả hàng"
	case Exception:
		return "Xử lý ngoài luồng"
	case Damage:
		return "Hàng hư hỏng"
	default:
		return "Trạng thái không xác định"
	}
}

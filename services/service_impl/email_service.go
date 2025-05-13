package serviceimpl

import (
	"703room/703room.com/services"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type emailService struct {
	ApiKey          string
	Expense_service services.UserHasPaymentService
	Room_service    services.RoomService
	User_service    services.UserService
}

func NewEmailService(apiKey string, expense_service services.UserHasPaymentService, room_service services.RoomService, user_service services.UserService) services.EmailService {
	return emailService{
		ApiKey:          apiKey,
		Expense_service: expense_service,
		Room_service:    room_service,
		User_service:    user_service,
	}
}

func (e emailService) SendReportToRoomate(ctx context.Context, room_id uuid.UUID, year string, month string, day string, data string) error {
	from := mail.NewEmail("RoomXpense88", "roomxpense88@gmail.com")
	log.Println(room_id)
	// Step 1: Get members in the room
	roomMembers, err := e.Room_service.ListMembersByRoomID(ctx, room_id.String())
	if err != nil {
		return err
	}
	log.Println(roomMembers)
	// Step 2: Build report content
	var reportContent string
	reportContent += "<h3>Báo cáo chi tiêu phòng</h3>"
	reportContent += "<table border='1' cellpadding='5' cellspacing='0'><tr><th>Họ tên</th><th>Email</th><th>Tổng chi tiêu</th></tr>"

	for _, member := range roomMembers {
		memberData, _ := e.User_service.GetUserByID(ctx, member.UserID.String())
		expense, err := e.Expense_service.CalculateMemberExpenseByMemberId(ctx, member.UserID, room_id, year, month, day)
		if err != nil {
			continue
		}
		reportContent += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%.2f VND</td></tr>", memberData.Name, memberData.Email, expense)
	}

	reportContent += "</table>"

	// Step 3: Compose subject and content
	subject := fmt.Sprintf("Báo cáo chi tiêu phòng - %s/%s", month, year)
	plainTextContent := "Vui lòng xem báo cáo chi tiêu"
	htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body>
			<p>Xin chào,</p>
			<p>Dưới đây là báo cáo chi tiêu cho phòng của bạn:</p>
			%s
			<p>Note của trưởng phòng: %s</p>
			<p>Trân trọng,<br>Hệ thống roomXpense</p>
		</body>
		</html>
	`, reportContent, data)

	// Step 4: Create a new V3Mail object
	message := mail.NewV3Mail()
	message.SetFrom(from)
	message.Subject = subject

	// Create a personalization object and add each member as a recipient
	personalization := mail.NewPersonalization()

	for _, member := range roomMembers {
		memberData, err := e.User_service.GetUserByID(ctx, member.UserID.String())
		if err != nil {
			continue
		}
		to := mail.NewEmail(memberData.Name, memberData.Email)
		personalization.AddTos(to)
	}

	message.AddPersonalizations(personalization)

	// Add content
	content := mail.NewContent("text/plain", plainTextContent)
	message.AddContent(content)
	contentHTML := mail.NewContent("text/html", htmlContent)
	message.AddContent(contentHTML)

	// Step 5: Send the email
	client := sendgrid.NewSendClient(e.ApiKey)
	log.Println(message)
	response, err := client.Send(message)
	log.Println("Enter line 98")
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("email send failed with status code: %d, body: %s", response.StatusCode, response.Body)
	}

	return nil
}

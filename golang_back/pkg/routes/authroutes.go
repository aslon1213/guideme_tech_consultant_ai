package routes

import (
	"aslon1213/customer_support_bot/pkg/handlers"
	"aslon1213/customer_support_bot/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(fb *fiber.App, md *middlewares.MiddlewaresWrapper, handlers *handlers.HandlersWrapper) {
	ah := handlers.AuthenticationHandlers
	//
	//

	keys := fb.Group("/keys", md.AuthenticationMiddleware)
	keys.Post("/create", ah.CreateKey)
	keys.Delete("/delete", ah.DeleteKey)
	keys.Get("/list", ah.ListTheKeys)
	keys.Get("/info", ah.GetInfo)

	auth := fb.Group("/auth")
	auth.Post("/login", ah.Login)
	auth.Post("/register", ah.Register)
	auth.Post("/logout", ah.Logout)
	auth.Post("/refresh", ah.Refresh)
	// fb.Post("/auth/forgot-password", ah.ForgotPassword)
	// fb.Post("/auth/reset-password", ah.ResetPassword)
	// fb.Post("/auth/verify", ah.Verify)
	// fb.Post("/auth/verify-email", ah.VerifyEmail)
	// fb.Post("/auth/verify-phone", ah.VerifyPhone)
	// fb.Post("/auth/verify-otp", ah.VerifyOTP)
	// fb.Post("/auth/verify-token", ah.VerifyToken)
	// fb.Post("/auth/verify-otp-token", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-email", ah.VerifyOTPWithToken)
	// fb.Post("/auth/verify-otp-phone", ah.VerifyOTPWithToken)
}

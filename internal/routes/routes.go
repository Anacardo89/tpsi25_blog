package routes

import (
	"github.com/Anacardo89/tpsi25_blog/internal/handlers/actions"
	"github.com/Anacardo89/tpsi25_blog/internal/handlers/pages"
	"github.com/Anacardo89/tpsi25_blog/internal/handlers/redirect"
	"github.com/Anacardo89/tpsi25_blog/internal/handlers/wsoc"
	"github.com/gorilla/mux"
)

func DeclareRoutes(r *mux.Router) {
	r.HandleFunc("/", redirect.RedirIndex).Schemes("https")
	r.HandleFunc("/home", pages.Home).Schemes("https")
	r.HandleFunc("/login", pages.Login).Schemes("https")
	r.HandleFunc("/register", pages.Register).Schemes("https")
	r.HandleFunc("/error", pages.Error).Schemes("https")
	r.HandleFunc("/newPost", pages.NewPost).Schemes("https")
	r.HandleFunc("/post/{post_guid}", pages.Post).Schemes("https")
	r.HandleFunc("/user/{encoded_user_name}", pages.UserProfile).Schemes("https")
	r.HandleFunc("/user/{encoded_user_name}/feed", pages.Feed).Schemes("https")
	r.HandleFunc("/user/{encoded_user_name}/followers", pages.UserFollowers).Schemes("https")
	r.HandleFunc("/user/{encoded_user_name}/following", pages.UserFollowing).Schemes("https")
	r.HandleFunc("/forgot-password", pages.ForgotPassword).Schemes("https")
	r.HandleFunc("/recover-password/{encoded_user_name}", pages.RecoverPassword).Schemes("https")
	r.HandleFunc("/change-password/{encoded_user_name}", pages.ChangePassword).Schemes("https")

	r.HandleFunc("/action/register", actions.RegisterUser).Methods("POST").Schemes("https")
	r.HandleFunc("/action/activate/{encoded_user_name}", actions.ActivateUser).Schemes("https")
	r.HandleFunc("/action/login", actions.Login).Methods("POST").Schemes("https")
	r.HandleFunc("/action/logout", actions.Logout).Methods("POST").Schemes("https")
	r.HandleFunc("/action/search/user", actions.SearchUsers).Methods("GET").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/follow", actions.RequestFollowUser).Methods("POST").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/accept", actions.AcceptFollowRequest).Methods("PUT").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/unfollow", actions.UnfollowUser).Methods("DELETE").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/profile-pic", actions.PostProfilePic).Methods("POST").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/conversations", actions.GetConversations).Methods("GET").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/conversations", actions.StartConversation).Methods("POST").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/conversations/{conversation_id}/read", actions.ReadConversation).Methods("PUT").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/conversations/{conversation_id}/dms", actions.GetDMs).Methods("GET").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/conversations/{conversation_id}/dms", actions.SendDM).Methods("POST").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/notifications", actions.GetNotifs).Methods("GET").Schemes("https")
	r.HandleFunc("/action/user/{encoded_user_name}/notifications/{notif_id}/read", actions.UpdateNotif).Methods("PUT").Schemes("https")
	r.HandleFunc("/action/post", actions.AddPost).Methods("POST").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}", actions.EditPost).Methods("PUT").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}", actions.DeletePost).Methods("DELETE").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}/up", actions.RatePostUp).Methods("POST").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}/down", actions.RatePostDown).Methods("POST").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}/comment", actions.AddComment).Methods("POST").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}/comment/{comment_id}", actions.EditComment).Methods("PUT").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}/comment/{comment_id}", actions.DeleteComment).Methods("DELETE").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}/comment/{comment_id}/up", actions.RateCommentUp).Methods("POST").Schemes("https")
	r.HandleFunc("/action/post/{post_guid}/comment/{comment_id}/down", actions.RateCommentDown).Methods("POST").Schemes("https")
	r.HandleFunc("/action/forgot-password", actions.ForgotPassword).Methods("POST").Schemes("https")
	r.HandleFunc("/action/recover-password", actions.RecoverPassword).Methods("POST").Schemes("https")
	r.HandleFunc("/action/change-password", actions.ChangePassword).Methods("POST").Schemes("https")

	r.HandleFunc("/ws", wsoc.HandleWebSocket)

	r.HandleFunc("/action/image", actions.PostImage).Schemes("https")
	r.HandleFunc("/action/profile-pic", actions.ProfilePic).Schemes("https")
}

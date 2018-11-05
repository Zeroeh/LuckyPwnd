public interface LuckyDayApi {
    @DELETE("/v3/Account/Delete")
    Call<Void> deleteAccount();

    @GET("/v3/User/EnteredLuckyCode")
    Call<EncryptedGsonObject> getEnteredLuckyCode();

    @GET("v3/Home/FreeTokens")
    Call<EncryptedGsonObject> getFreeTokens();

    @GET("/v3/Games/Leaderboard")
    Call<EncryptedGsonObject> getLeaderboard();

    @GET("/v3/Winners/{page}/{rewardType}")
    Call<EncryptedGsonObject> getLuckyWinners(@Path("page") int i, @Path("rewardType") int i2);

    @GET("/v3/User/History")
    Call<EncryptedGsonObject> getOrderHistory(@Query("page") int i);

    @GET("/v3/User/Account")
    Call<EncryptedGsonObject> getProfile();

    @GET("/v3/RaffleGames")
    Call<EncryptedGsonObject> getRaffleGames();

    @GET("/v3/RaffleGames/GetRaffleById")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> getRaffleGamesById(@Query("id") int i);

    @GET("/v3/Home/IsRateDisplayed")
    Call<EncryptedGsonObject> getRateDisplayed();

    @POST("/v3/Games/Scratcher/WelcomeScratcher")
    Call<EncryptedGsonObject> getWelcomeScratcher();

    @POST("/v3/User/ApplyLuckyCode")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postApplyLuckyCode();

    @POST("/v3/RaffleGames/Purchase")
    Call<EncryptedGsonObject> postBuyRaffleTickets(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/User/ChangePassword")
    Call<EncryptedGsonObject> postChangePassword(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/BlackJack/Complete")
    Call<EncryptedGsonObject> postCompleteBlackJackGame(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/FacebookLike")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postFBLike();

    @POST("/v3/Games/Feedback")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postFeedback(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Account/ForgotPassword")
    Call<EncryptedGsonObject> postForgotPassword(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/FortuneWheel")
    Call<EncryptedGsonObject> postFortuneWheel();

    @POST("/v3/FreeTokensCampaigns/Grant")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postGrants(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Home/UpdateHomePage")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postHomePageInfo(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/InstantRewards/Details")
    Call<EncryptedGsonObject> postInstantRewardDetails(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/InstantRewards")
    Call<EncryptedGsonObject> postInstantRewards(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/BlackJack/Insure")
    Call<EncryptedGsonObject> postInsureBlackJackGame(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/Lotto/Details")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postLottoDetails(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Payment/Withdrawal")
    Call<EncryptedGsonObject> postMoneyWithdrawal(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Notifications/Token")
    Call<EncryptedGsonObject> postNotificationToken(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/offerwall/reward")
    Call<EncryptedGsonObject> postOfferwallReward();

    @POST("/v3/InstantRewards/Refund")
    Call<EncryptedGsonObject> postOrderHistoryRefund(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/User/Account")
    Call<Void> postProfileUpdate(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/InstantRewards/PurchaseGift")
    Call<EncryptedGsonObject> postPurchaseGift(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/InstantRewards/PurchaseGiftCard")
    Call<EncryptedGsonObject> postPurchaseGiftCard(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/RaffleGames/Purchase")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postRaffleGamesPurchase(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/RaffleGames/TicketOdds")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postRaffleGamesTicketOdds(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/RaffleGames/RevealTickets")
    Call<EncryptedGsonObject> postRevealRaffleGamesTickets(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/Lotto/RevealWin")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postRevealWin(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/User/SaveLuckyCode")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> postSaveLuckyCode(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/Scratcher/ScratcherCheckpoint")
    Call<EncryptedGsonObject> postScratcherCheckpoint();

    @POST("/v3/Games/Scratcher/ScratcherDetails")
    Call<EncryptedGsonObject> postScratcherDetails(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/Scratcher/Play")
    Call<EncryptedGsonObject> postScratcherGamePlay(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Share")
    Call<EncryptedGsonObject> postSharesEvent(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/Scratcher/ScratchersSocialActivity")
    Call<Void> postSocialScratcherActivity(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/BlackJack/Start")
    Call<EncryptedGsonObject> postStartBlackJackGame(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Games/Lotto/StartGame")
    Call<EncryptedGsonObject> postStartLottoGame(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/User/Avatar")
    @Multipart
    Call<EncryptedGsonObject> postUploadAvatar(@Part MultipartBody$Part multipartBody$Part);

    @POST("/v3/User/SaveLuckyImage")
    @Multipart
    Call<EncryptedGsonObject> postUploadLuckyImage(@Part MultipartBody$Part multipartBody$Part);

    @POST("/v3/User/SaveShareAvatar")
    @Multipart
    Call<EncryptedGsonObject> postUploadShare(@Part("PrizeType") TypedString typedString, @Part("Amount") TypedString typedString2, @Part("Type") TypedString typedString3, @Part MultipartBody$Part multipartBody$Part);

    @POST("/v3/Account/VerifyEmail")
    Call<EncryptedGsonObject> postVerifyEmail(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/FreeChips/Videos")
    Call<EncryptedGsonObject> postVideoFreeCredits(@Body EncryptedGsonObject encryptedGsonObject);
}

public interface LuckyDayHttpsApi {
    @POST("/v3/Account/LoginEmail")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> loginEmail(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Account/LoginFacebook")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> loginFacebook(@Body EncryptedGsonObject encryptedGsonObject);

    @POST("/v3/Account/Logout")
    @Headers({"api-version:3"})
    Call<Void> logout();

    @POST("/v3/Account/Registration")
    @Headers({"api-version:3"})
    Call<EncryptedGsonObject> registerUser(@Body EncryptedGsonObject encryptedGsonObject);
}


public interface ApiMaps {
    public static final String END_POINT = "https://maps.googleapis.com";

    @GET("/maps/api/geocode/json")
    Call<JsonElement> getInfoByCoord(@Query("latlng") String str, @Query("language") String str2, @Query("key") String str3);

    @GET("/maps/api/place/details/json")
    Call<JsonElement> getPlaceDetails(@Query("key") String str, @Query("placeid") String str2);
}








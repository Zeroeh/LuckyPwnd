# LuckyPwnd
A clientless bot for an app called "Lucky Day"
The app gives out cash prizes for playing scratchers and entering daily lottery drawings.

I had originally created this so that I could see if the app was vulnerable to having tens of thousands of accounts claiming free money from it, but they quickly adapted to my techniques and so I gave up. I'm sure one could get away by using only a couple hundred accounts and using individual proxies for each account. It also appears that all the logic is server sided with no way for the user to guarentee that they will earn any cash at all. What I found out through my analysis is that you are eligible for ~$1 on the first round of scratchers but after that the system no longer authenticates any more cash prizes for the user and so it seems like the devs of lucky day can block winnings as well as control how much cash is given out per user.


I hastily made this app and had no long term plans for this project and thus the code is a bit... all over the place. A lot of the core abstract behavior will need to be refactored/redone entirely. The progressive logic flow of the bot should still be easy to decipher however. Most of the juicy stuff is happening in the bot.go and client.go files. Most of the bot logic and behavior is in bot.go. I redacted information from files that could trace back to me. While redacting info, some of the functions may have broke and the entire app could not work as a result. I haven't tested and I don't care. I'm just uploading this for educational use.

I will not help in setting this up, neither will I accept any issues or pull requests.

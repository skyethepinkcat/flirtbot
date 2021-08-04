package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	flirt := "872555017352126546";
	noflirt := "872555070200373258";
	guildid := m.GuildID;
	userid := m.Author.ID;
	member, _ := s.GuildMember(guildid, userid)
	// if contains(member.Roles, flirt) {
	// If the message is "ping" reply with "Pong!"
	if strings.HasPrefix(m.Content,"!flirt") {
		if contains(member.Roles, flirt) {
			fmt.Println(s.GuildMemberRoleAdd(guildid, userid, noflirt));
			s.ChannelMessageSend(m.ChannelID, "Changed your role so people **won't** flirt with you~")
		} else {
			fmt.Println(s.GuildMemberRoleAdd(guildid, userid, flirt));
			s.ChannelMessageSend(m.ChannelID, "Changed your role so people **will** flirt with you~")
		}
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

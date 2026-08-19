# Learning & Blocker Journal — Assignment 1

**Tool assigned:** Message Queues
**Time-box:** 4 hours
**Start time:** ~12:00 AM
**End time:** ~11:00 AM (note: included a sleep break overnight — active work time was well within the 4-hour box, but wall-clock time spanned ~11 hours due to pausing to rest)

---

## 1. Orientation Notes (fill in as you research)
- **What is this tool, in your own words?** A message queue lets one program send a message without waiting for another program to be ready to receive it — the message sits in the queue until a consumer picks it up.
- **What problem does it solve that you couldn't easily solve without it?** Without a queue, two systems calling each other directly have to both be online and responsive at the same time, and one can get stuck waiting on the other. A queue decouples them — the sender doesn't need the receiver to be available right now.
- **What's the simplest possible use case you could build?** A producer publishing a message to a subject, and a consumer subscribed to that same subject printing it out when received — basic pub/sub.

---

## 2. Resources Consulted
| Time | Resource (link/doc/video) | Why you went there | What you got from it |
|------|---------------------------|---------------------|------------------------|
| ~1-2 AM | docs.nats.io (official NATS docs) | To find setup/quickstart instructions for the server | Confirmed the real docs site (after initially landing on a Synadia marketing/workshop page by mistake); pointed toward GitHub for the actual server and client code |
| ~2 AM | github.com/nats-io/nats-server and github.com/nats-io/nats.go | To get the server binary and the Go client library | Found the Releases page for the server (used `go install` instead since Windows binaries weren't obviously listed) and the nats.go README with basic pub/sub code examples |

---

## 3. Blocker Log (the most important section — fill this in REAL TIME, not after)
For every single thing that goes wrong, no matter how small:

### Blocker #1
- **Time hit:** ~12-1 AM by running `go version` in VS Code's PowerShell terminal
- **Exact error / symptom:** `go : The term 'go' is not recognized as the name of a cmdlet, function, script file, or operable program.` Also tried `Go version` (capital G) and got a similar "not recognized" error. Then tried `"go version"` (in quotes) which just echoed the text back instead of running anything.
- **What I tried (in order):** 
  1. Ran `go version` — failed, command not recognized
  2. Ran `Go version` (capitalized) — also failed, PowerShell treated it as a different/unknown command entirely
  3. Ran `"go version"` in quotes — PowerShell just printed the literal string back, since quotes make it text, not a command
- **What actually fixed it:** Fully closing VS Code (not just the terminal tab) and reopening it. The terminal had been opened before PATH was updated with Go's install location, so it was running on a stale PATH. A full restart forced it to pick up the change. Confirmed working at 1:06 AM: `go version go1.26.6 windows/amd64`
- **Time lost:** ~10-30 min
- **What I'd do differently next time:** Check the Environment Variables PATH list directly to confirm Go's install location (`C:\Program Files\Go\bin`) is present before assuming the terminal is the problem. Also: PowerShell commands are case-sensitive, so `Go` ≠ `go`.

### Blocker #2
- **Time hit:** ~9:00 AM
- **What I was trying to do:** Keep subscriber.go running so it wouldn't exit immediately after subscribing, using `select {}` at the end of main()
- **Exact error / symptom:** `fatal error: all goroutines are asleep - deadlock! goroutine 1 [select (no cases)]: main.main() ... exit status 2`
- **What I tried (in order):**
  1. Used `select {}` (empty braces) expecting it to block/wait indefinitely
- **What actually fixed it:** Empty `select {}` has zero cases, so Go's runtime detects it can never proceed and kills the program as a deadlock rather than hanging forever. Replaced with `time.Sleep(time.Hour)` (and added "time" to imports) to keep the program alive long enough to receive messages for a prototype/testing purposes.
- **Time lost:** ~10-30 min
- **What I'd do differently next time:** Understand that `select {}` needs actual cases (channel operations) to be a legitimate "wait forever" pattern — an empty one is a no-op deadlock, not a blocking wait. For a quick prototype, `time.Sleep()` is a simpler substitute; for production code, a real channel-based block or signal handling would be more correct.

### Blocker #3
- **Time hit:** ~10:00 AM
- **What I was trying to do:** Debug why subscriber.go never printed "Received a message" even though publisher.go ran successfully every time with no errors
- **Exact error / symptom:** No error at all — publisher printed "message sent" consistently, subscriber just sat silently with no output. Misleading because there was no crash to point to the real cause.
- **What I tried (in order):**
  1. Checked if I was typing commands into the wrong terminal tab — ruled out, publisher was genuinely running and printing success
  2. Asked to check the nats-server.exe terminal window for connection logs
  3. Discovered that terminal no longer existed — it had been running inside a VS Code integrated terminal, and closing/switching the VS Code workspace folder (from BlackJack to nats-prototype) silently closed that terminal, killing the server process
- **What actually fixed it:** Restarted nats-server.exe in a standalone PowerShell window OUTSIDE of VS Code entirely, so it's immune to VS Code workspace/folder switches. Only closing that specific window or Ctrl+C inside it will stop the server going forward.
- **Time lost:** ~30-60 min
- **What I'd do differently next time:** Recognize that VS Code integrated terminals are tied to the workspace session — long-running background processes (like a server) should run in a terminal outside the IDE, not inside it, if the IDE workspace might change. Also: `nc, _ := nats.Connect(...)` silently discarding the error with `_` hid the real problem — should check that error properly (e.g. `if err != nil { log.Fatal(err) }`) so a failed connection is loud, not silent.

### Blocker #4
(copy the block above for each new blocker)

---

## 4. Time-to-Completion Tracking
**Note on timing:** Work spanned ~12:00 AM to ~11:00 AM (~11 hours wall-clock), but this included an overnight sleep break. "Time actually taken" below reflects active, hands-on-keyboard work only — not the sleep gap.

| Phase | Time budgeted | Active time taken | Notes |
|-------|---------------|----------------------|-------|
| Orientation | ~45 min | ~30-45 min | Found docs.nats.io, initially landed on unrelated Synadia marketing page first |
| Setup | ~30 min | ~1-1.5 hrs | Included Blocker #1 (go PATH issue, ~12-1 AM) |
| Build | ~1.5-3 hrs | ~1-1.5 hrs | Wrote publisher.go and subscriber.go using nats.go README examples; included Blocker #2 (select{} deadlock, ~9 AM) |
| Stress test | ~40 min | ~30-60 min | Blocker #3 (server terminal silently closed, ~10 AM) doubled as an unplanned stress test of a real failure mode |
| Wrap-up | ~20 min | ~20-30 min | Journal cleanup and reflection |
| **Total active time** | ~4 hrs | ~3.5-4.5 hrs | Roughly matched the original 4-hour budget |
| **Total wall-clock time** | — | ~11 hrs (12 AM–11 AM) | Includes an overnight sleep break between the setup/build phase and the stress-test/wrap-up phase |

---

## 5. Final Reflection
- **What does your prototype actually do (2-3 sentences)?** A publisher connects to a locally running NATS server and publishes a "Hello World" message to a subject called "Orders." A separate subscriber process connects to the same server, listens on that subject, and prints the message when it arrives. Together they demonstrate the core producer → queue/broker → consumer pattern of a message queue.
- **What's still broken or missing?** Error handling is minimal — the connection error from `nats.Connect()` is discarded with `_` instead of being checked, so failures (like the server being down) fail silently instead of giving a clear error message. The subscriber also uses `time.Sleep(time.Hour)` as a simple placeholder to stay alive rather than a more proper blocking/signal-based approach.
- **What's the single biggest thing you learned about troubleshooting alone (not about the tool itself)?** Understand the basics of message queues now, though details are still fuzzy. The biggest source of friction wasn't the message queue concept itself — it was fighting tool and environment setup issues (PATH, terminal sessions tied to the IDE, silent failures with no error message). Learned that environment/tooling problems are often harder to diagnose than code bugs because they don't throw clear errors; troubleshooting them required checking assumptions methodically (which terminal, is the process still running, what closed it) rather than reading a stack trace.
- **If you had 4 more hours, what would you do next?** Would try a different NATS pattern beyond basic pub/sub — like a synchronous subscriber or the request-reply pattern shown in the docs — to understand more of what the tool can do beyond the minimum working example, and add proper error handling instead of discarding errors with `_`.

---

## Next Steps (paused here — resume from this point)

**Tool chosen:** NATS (no Docker needed — runs as a standalone .exe on Windows)
**Environment confirmed working:** Go installed, `go version` returns `go1.26.6 windows/amd64`

1. Go to the official NATS docs — search "NATS quickstart Go" or go directly to nats.io's documentation site. Look for "Getting Started" or "Quickstart."
2. Download and run the NATS server — a single `.exe`, no installer, no Docker. Running it starts a local server listening on a port (commonly `localhost:4222` by default — confirm in docs).
3. Get the Go client library for NATS — likely via a `go get` command shown directly on the docs page. This lets your Go code actually talk to the NATS protocol.

**Target shape of the prototype:**
- One small Go file that connects to NATS and **publishes** ("produces") a message
- One small Go file that connects to NATS and **subscribes** ("consumes") and prints what it receives
- That's the minimum viable prototype — no more needed to satisfy Assignment 1

---

### Progress Log — publisher.go
- Wrote a complete publisher using the nats.go README example: connected via `nats.Connect(nats.DefaultURL)`, published a message to subject "Orders" with `nc.Publish(...)`.
- Ran with `go run publisher.go` — succeeded on first try, printed "message sent," no errors.
- Note: since no subscriber was running yet, the published message wasn't received by anything — NATS core pub/sub doesn't queue messages for subscribers that join late (unlike some other queue systems). Confirmed this by observing no receiver output.
- Currently building subscriber.go to complete the pub/sub loop.

### Progress Log — subscriber.go
- Fixed the select{} deadlock (see Blocker #2), subscriber ran successfully.
- Saw exit code `0xc000013a` when stopping it — initially unclear if this was an error. Investigated: this is Windows' STATUS_CONTROL_C_EXIT code, confirming the program was manually interrupted via Ctrl+C, not a crash or bug. Expected behavior, no fix needed.

### Progress Log — Full pub/sub loop working
- Root cause of Blocker #3 confirmed and fixed: server was down because its terminal got closed with the VS Code workspace switch. Restarted nats-server.exe in a standalone PowerShell window, independent of VS Code.
- Ran subscriber.go (waiting/listening), then publisher.go in a separate terminal.
- CONFIRMED WORKING: subscriber terminal printed "Received a message: Hello World" after publisher ran. Server window showed new connection activity for both processes.
- This is the core functional prototype for Assignment 1: a producer publishes to a subject ("Orders"), the NATS server routes it, and a consumer receives and prints it — the minimum viable message queue demonstration.

## 6. Prototype Summary
- **Repo/file location:** C:\Users\user\nats-prototype (publisher.go, subscriber.go, go.mod, go.sum)
- **How to run it:** 
  1. Start the NATS server: `C:\Users\user\go\bin\nats-server.exe` (run in a standalone terminal, NOT inside VS Code, so it survives workspace switches)
  2. In a separate terminal, run the subscriber: `go run subscriber.go` (leave running — it listens for 1 hour via time.Sleep)
  3. In another terminal, run the publisher: `go run publisher.go`
- **What "working" looks like when you run it:** Publisher prints "message sent." Subscriber terminal prints "Received a message: Hello World." Server terminal shows connection activity for both.

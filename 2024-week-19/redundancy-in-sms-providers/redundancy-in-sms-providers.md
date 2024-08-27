# Redundancy in SMS Providers

- Agenda

  - Usage scenario
  - Previous failures
  - Introduce a small crisis
  - Code snippets
  - Stats
  - Possibilities
  - Money

- Usage scenario
  - When do people log in?
  - In the shopping situation.
  - Personally, I log in when in line for the cash register
  - Worst situation to have something go wrong
  - My toddler has been promised a yoghurt and is very anxious to get it
  - Other people are waiting in line behind me
- 99 problems but passwords ain't one
- SMS is out of my control
- External vendor for SMS, sort of out of the team's control as well

- The IDP team sends SMSs as a multifactor authentication method
  - Our provider is LinkMobility
  - They had downtime once, and as a result, so did we
  - Avoid downtime by rendundancy
  - Demo [login at coop.no](https://www.coop.no/)

Previous failures

- Usually it's fine. More than 99,7 % of the time, it's fine.
- 99,7 % translates to 5 minutes a day (avg)
- When we have downtimes, that number looks different
- Jun 4, 8:00 am – Jun 4, 11:11 am
  - [Acceptance](https://app.datadoghq.eu/dashboard/aqp-5na-cwj?fromUser=true&refresh_mode=paused&view=spans&from_ts=1717480800000&to_ts=1717492260000&live=false)
- Jul 16, 10:59 am – Jul 16, 12:39 pm

  - [Acceptance](https://app.datadoghq.eu/dashboard/aqp-5na-cwj?fromUser=true&refresh_mode=paused&view=spans&from_ts=1721120340000&to_ts=1721126340000&live=false)

- Introduce a small crisis

  - How to demo a redundancy
  - It should always just work
  - [Let's take down production](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*ylTf2MZ7GSRHfHv04-ccWA.jpeg)

-Code snippets

- Let's look at [how it works](https://github.com/coopnorge/idp-auth0-config/tree/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage):

- [The control flow](https://github.com/coopnorge/idp-auth0-config/blob/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage/sms.ts)
- [The data](https://github.com/coopnorge/idp-auth0-config/blob/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage/smsTypes.ts)
- [An SMS provider](https://github.com/coopnorge/idp-auth0-config/blob/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage/linkMobility.ts)
- [Stats](https://github.com/coopnorge/idp-auth0-config/blob/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage/logResults.ts)

- Look at the live stats

  - Revisit stats to inspect that we
    - started failing sending smss with our primary provider
    - the secondary provider picked up the slack

- Types of errors

  - Immediate
  - Delayed

- Benefits as of now

  - Built-in incident handling
  - Bonus: Transient error handling
  - Bonus: Trivial to switch providers over non-error issues
    - price

- How to take this forward?

  - Purposes
    - Handle delayed errors
    - Serve other teams
    - Build out to other types of messages (email?)
  - Function -> server
  - Storing generated messages with internal ids for re-referencing

- Money
  - Switching main providers to a cheaper options has saved the IDP
  - 338.000,- yearly
  - Other teams send more SMSs than we do

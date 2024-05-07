- The IDP team send SMSs as a multifactor authentication method

  - Our provider is LinkMobility
  - They had downtime once, and as a result, so did we
  - Avoid downtime by rendundancy

- How to demo a redundancy

  - It should always just work
  - [Let's take down production](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*ylTf2MZ7GSRHfHv04-ccWA.jpeg)

- Let's look at [how it works](https://github.com/coopnorge/idp-auth0-config/tree/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage):

  - [The control flow](https://github.com/coopnorge/idp-auth0-config/blob/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage/sms.ts)
  - [The data](https://github.com/coopnorge/idp-auth0-config/blob/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage/smsTypes.ts)
  - [An SMS provider](https://github.com/coopnorge/idp-auth0-config/blob/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage/linkMobility.ts)
  - [Stats](https://github.com/coopnorge/idp-auth0-config/blob/0ea100e4010b73779719d7a8605354c2549e0820/scripts/src/actions/sendPhoneMessage/logResults.ts)

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

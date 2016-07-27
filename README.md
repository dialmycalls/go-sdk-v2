# Go API client for dialmycalls

The DialMyCalls API

## Overview

For more information, please visit [https://www.dialmycalls.com](https://www.dialmycalls.com)

## Installation
Put the package under your project folder and add the following in import:
```
    "./dialmycalls"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.dialmycalls.com/2.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Accounts* | [**CreateAccessAccount**](docs/Accounts.md#createaccessaccount) | **Post** /accessaccount | Add Access Account
*Accounts* | [**DeleteAccessAccountById**](docs/Accounts.md#deleteaccessaccountbyid) | **Delete** /accessaccount/{AccessAccountId} | Delete Access Account
*Accounts* | [**GetAccessAccountById**](docs/Accounts.md#getaccessaccountbyid) | **Get** /accessaccount/{AccessAccountId} | Get Access Account
*Accounts* | [**GetAccessAccounts**](docs/Accounts.md#getaccessaccounts) | **Get** /accessaccounts | List Access Accounts
*Accounts* | [**GetAccount**](docs/Accounts.md#getaccount) | **Get** /account | Get Account
*Accounts* | [**UpdateAccessAccountById**](docs/Accounts.md#updateaccessaccountbyid) | **Put** /accessaccount/{AccessAccountId} | Update Access Account
*CallerIds* | [**CreateCallerId**](docs/CallerIds.md#createcallerid) | **Post** /callerid | Add Caller ID
*CallerIds* | [**CreateUnverifiedCallerId**](docs/CallerIds.md#createunverifiedcallerid) | **Post** /verify/callerid | Add Caller ID (Unverified)
*CallerIds* | [**DeleteCallerIdById**](docs/CallerIds.md#deletecalleridbyid) | **Delete** /callerid/{CalleridId} | Delete Caller ID
*CallerIds* | [**GetCallerIdById**](docs/CallerIds.md#getcalleridbyid) | **Get** /callerid/{CalleridId} | Get Caller ID
*CallerIds* | [**GetCallerIds**](docs/CallerIds.md#getcallerids) | **Get** /callerids | List Caller IDs
*CallerIds* | [**UpdateCallerIdById**](docs/CallerIds.md#updatecalleridbyid) | **Put** /callerid/{CalleridId} | Update Caller ID
*CallerIds* | [**VerifyCallerIdById**](docs/CallerIds.md#verifycalleridbyid) | **Put** /verify/callerid/{CalleridId} | Verify Caller ID
*Calls* | [**CancelCallById**](docs/Calls.md#cancelcallbyid) | **Delete** /service/call/{CallId} | Cancel Call
*Calls* | [**CreateCall**](docs/Calls.md#createcall) | **Post** /service/call | Create Call
*Calls* | [**GetCallById**](docs/Calls.md#getcallbyid) | **Get** /service/call/{CallId} | Get Call
*Calls* | [**GetCallRecipientsByCallId**](docs/Calls.md#getcallrecipientsbycallid) | **Get** /service/call/{CallId}/recipients | Get Call Recipients
*Calls* | [**GetCalls**](docs/Calls.md#getcalls) | **Get** /service/calls | List Calls
*Contacts* | [**CreateContact**](docs/Contacts.md#createcontact) | **Post** /contact | Add Contact
*Contacts* | [**DeleteContactById**](docs/Contacts.md#deletecontactbyid) | **Delete** /contact/{ContactId} | Delete Contact
*Contacts* | [**GetContactById**](docs/Contacts.md#getcontactbyid) | **Get** /contact/{ContactId} | Get Contact
*Contacts* | [**GetContacts**](docs/Contacts.md#getcontacts) | **Get** /contacts | List Contacts
*Contacts* | [**GetContactsByGroupId**](docs/Contacts.md#getcontactsbygroupid) | **Get** /contacts/{GroupId} | List Contacts in Group
*Contacts* | [**UpdateContactById**](docs/Contacts.md#updatecontactbyid) | **Put** /contact/{ContactId} | Update Contact
*DoNotContacts* | [**GetDoNotContacts**](docs/DoNotContacts.md#getdonotcontacts) | **Get** /donotcontacts | List DoNotContacts
*Groups* | [**CreateGroup**](docs/Groups.md#creategroup) | **Post** /group | Add Group
*Groups* | [**DeleteGroupById**](docs/Groups.md#deletegroupbyid) | **Delete** /group/{GroupId} | Delete Group
*Groups* | [**GetGroupById**](docs/Groups.md#getgroupbyid) | **Get** /group/{GroupId} | Get Group
*Groups* | [**GetGroups**](docs/Groups.md#getgroups) | **Get** /groups | List Groups
*Groups* | [**UpdateGroupById**](docs/Groups.md#updategroupbyid) | **Put** /group/{GroupId} | Update Group
*Keywords* | [**DeleteKeywordByID**](docs/Keywords.md#deletekeywordbyid) | **Delete** /keyword/{KeywordId} | Delete Keyword
*Keywords* | [**GetKeywordByID**](docs/Keywords.md#getkeywordbyid) | **Get** /keyword/{KeywordId} | Get Keyword
*Keywords* | [**GetKeywords**](docs/Keywords.md#getkeywords) | **Get** /keywords | List Keywords
*Recordings* | [**CreateRecording**](docs/Recordings.md#createrecording) | **Post** /recording/tts | Create Recording (Text-to-Speech)
*Recordings* | [**CreateRecordingByPhone**](docs/Recordings.md#createrecordingbyphone) | **Post** /recording/phone | Create Recording (Phone)
*Recordings* | [**CreateRecordingByUrl**](docs/Recordings.md#createrecordingbyurl) | **Post** /recording/url | Create Recording (URL)
*Recordings* | [**DeleteRecordingById**](docs/Recordings.md#deleterecordingbyid) | **Delete** /recording/{RecordingId} | Delete Recording
*Recordings* | [**GetRecordingById**](docs/Recordings.md#getrecordingbyid) | **Get** /recording/{RecordingId} | Get Recording
*Recordings* | [**GetRecordings**](docs/Recordings.md#getrecordings) | **Get** /recordings | List Recordings
*Recordings* | [**UpdateRecordingById**](docs/Recordings.md#updaterecordingbyid) | **Put** /recording/{RecordingId} | Update Recording
*Texts* | [**CancelTextById**](docs/Texts.md#canceltextbyid) | **Delete** /service/text/{TextId} | Cancel Text
*Texts* | [**CreateText**](docs/Texts.md#createtext) | **Post** /service/text | Create Text
*Texts* | [**DeleteIncomingTextById**](docs/Texts.md#deleteincomingtextbyid) | **Delete** /incoming/text/{TextId} | Delete Incoming Text
*Texts* | [**GetIncomingTextById**](docs/Texts.md#getincomingtextbyid) | **Get** /incoming/text/{TextId} | Get Incoming Text
*Texts* | [**GetIncomingTexts**](docs/Texts.md#getincomingtexts) | **Get** /incoming/texts | List Incoming Texts
*Texts* | [**GetShortCodes**](docs/Texts.md#getshortcodes) | **Get** /shortcodes | List Shortcodes
*Texts* | [**GetTextById**](docs/Texts.md#gettextbyid) | **Get** /service/text/{TextId} | Get Text
*Texts* | [**GetTextRecipientsByTextId**](docs/Texts.md#gettextrecipientsbytextid) | **Get** /service/text/{TextId}/recipients | Get Text Recipients
*Texts* | [**GetTexts**](docs/Texts.md#gettexts) | **Get** /service/texts | List Texts
*VanityNumbers* | [**DeleteVanityNumberById**](docs/VanityNumbers.md#deletevanitynumberbyid) | **Delete** /vanitynumber/{VanityNumberId} | Delete Vanity Number
*VanityNumbers* | [**GetVanityNumberById**](docs/VanityNumbers.md#getvanitynumberbyid) | **Get** /vanitynumber/{VanityNumberId} | Get Vanity Number
*VanityNumbers* | [**GetVanityNumbers**](docs/VanityNumbers.md#getvanitynumbers) | **Get** /vanitynumbers | List Vanity Numbers
*VanityNumbers* | [**UpdateVanityNumberById**](docs/VanityNumbers.md#updatevanitynumberbyid) | **Put** /vanitynumber/{VanityNumberId} | Update Vanity Number


## Documentation For Models

 - [Accessaccount](docs/Accessaccount.md)
 - [Account](docs/Account.md)
 - [CallRecipient](docs/CallRecipient.md)
 - [Callerid](docs/Callerid.md)
 - [Callservice](docs/Callservice.md)
 - [Contact](docs/Contact.md)
 - [ContactAttributes](docs/ContactAttributes.md)
 - [CreateAccessAccountParameters](docs/CreateAccessAccountParameters.md)
 - [CreateCallParameters](docs/CreateCallParameters.md)
 - [CreateCallerIdParameters](docs/CreateCallerIdParameters.md)
 - [CreateContactParameters](docs/CreateContactParameters.md)
 - [CreateGroupParameters](docs/CreateGroupParameters.md)
 - [CreateRecordingByPhoneParameters](docs/CreateRecordingByPhoneParameters.md)
 - [CreateRecordingByUrlParameters](docs/CreateRecordingByUrlParameters.md)
 - [CreateRecordingParameters](docs/CreateRecordingParameters.md)
 - [CreateTextParameters](docs/CreateTextParameters.md)
 - [CreateUnverifiedCallerIdParameters](docs/CreateUnverifiedCallerIdParameters.md)
 - [Donotcontact](docs/Donotcontact.md)
 - [Group](docs/Group.md)
 - [Identifier](docs/Identifier.md)
 - [Incomingtext](docs/Incomingtext.md)
 - [Keyword](docs/Keyword.md)
 - [Polling](docs/Polling.md)
 - [PushToListenAgain](docs/PushToListenAgain.md)
 - [PushToOptOut](docs/PushToOptOut.md)
 - [PushToRecord](docs/PushToRecord.md)
 - [PushToTalk](docs/PushToTalk.md)
 - [Recording](docs/Recording.md)
 - [Service](docs/Service.md)
 - [Shortcode](docs/Shortcode.md)
 - [TextRecipient](docs/TextRecipient.md)
 - [UpdateAccessAccountByIdParameters](docs/UpdateAccessAccountByIdParameters.md)
 - [UpdateCallerIdByIdParameters](docs/UpdateCallerIdByIdParameters.md)
 - [UpdateContactByIdParameters](docs/UpdateContactByIdParameters.md)
 - [UpdateGroupByIdParameters](docs/UpdateGroupByIdParameters.md)
 - [UpdateRecordingByIdParameters](docs/UpdateRecordingByIdParameters.md)
 - [UpdateVanityNumberByIdParameters](docs/UpdateVanityNumberByIdParameters.md)
 - [Vanitynumber](docs/Vanitynumber.md)
 - [VerifyCallerIdByIdParameters](docs/VerifyCallerIdByIdParameters.md)


## Documentation For Authorization


## api_key

- **Type**: API key 
- **API key parameter name**: X-Auth-ApiKey
- **Location**: HTTP header


## Author

support@dialmycalls.com


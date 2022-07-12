# What OpenBillsAPI is supposed to be

An API that receives HTTP requests to generate invoices and.. stuff

Currently I'm only working on the backend part.

## Some goals I must achieve

1. This API should use websockets to handle storing temporary data on server side.
    Reasons for that:
    - If the OpenBills client crashes the server still has the data "autosaved", basically
    - It's easier to create a client that way I guess, like I want to have seperate desktop and mobile clients 
      (mobile is completely tentative) so I don't have to worry many things since the Go backend would handle some boring stuff.

    Problems with that:
    - Probably is very slow if the server is hosted on the internet (i.e on a VPS which is how at least I'd use it)


## Some goals I'd (probably) like to achieve

1. Have an option to turn off calculation of boring stuff on server-side so if the network is slow basic 
   functionality isn't slowed down (i.e only the database is functioning over the network)

2. Being able to *configure* most things. I don't know what features will make it into OpenBills but 
   I want it to be customizable so everyone can benifit from it without much programming knowledge.  
   The config part will not be something user-friendly. You might still need linux knowledge for that. 
   (also this is just a hypothetical "feature" who knows what way it goes)


## What is a client?

OpenBills is just a server. It handles database functions and authentication, *caching* and similar stuff that the client side doesn't really need to do.
It's a bit of an unnecessary feature but I'd probably also make it do calculations and handling temporary data 
(i.e if you're generating an invoice it'd store the currently added items on the server side and additons/deletions occour on the server side. Optionally.)  

## Who OpenBills is for

Currently I do not plan to adapt OpenBills to fit "everyone's needs". I'm creating this for someone I know and I hope that I can create a good quality product 
that even others would like to use. But that is NOT currently my focus. Because I've killed this project before so I want to do things my way at the start..

OpenBills is free software, so you can definitely add features or help me develop this project (please do it!!), but I obviously can't gurantee the changes 
would make it into OpenBills. Also, OpenBills will be GST based and there are no plans to support other taxation types since GST is the one used where I live and
really that's what I care about.

## What's the point of this document

This is sort of like my "public TODO list"

Others can see this to understand where the project plans to go, and maybe even suggest some things!

Also seems like if I keep my thoughts public I'd be more motivated.

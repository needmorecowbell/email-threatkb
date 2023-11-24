# EML-Funnel Worker

This worker consists of
- an email relay which sends the raw eml to be processed by the server included in the root folder of this repository
- a collection of rest api endpoints for managing the routing configurations,stored in KV.

## Goal

The goal of this worker is to act as a quick way to generate a burner email that gets forwarded to your personal email only once it's been checked. If the message is malicious, it is forwarded to a separate location for investigation. Since this is a catch all address and I can find the recipient address via this worker, I'd like to also extend this as a routing feature for anyone else who would like to be protected by my detections. Long term, it would be cool to have a dashboard that would show the stats for each email -- how many senders have sent to this address, how many of those messages were malicious, graphs of the activity to the email over time.. etc. Would also like the option to forward to null and just collect the email if it's been detected as bad.


## Usage

Setup your kv namespace, set up the server and configure your dns to point at it's location for the subdomain of your choice (ie processor.domain.com), using a tunnel or by other means. You'll also want to set at least 2 custom addresses in Email -> Email Routing, one for the vault and one for the passthrough. You can have the vault and the passthrough address forward to the same destination. With those set up, you'll then configure your wrangler.toml with the information you've now set up. With that done, deploy the worker. You'll now need to associate this worker to the catch-all address. Go back to the Email Routing and add the catch-all, actioning it to send to a worker, that worker name being eml-funnl.

If it doesn't show up you may need to associate first it in the Email -> Email Workers tab.. not sure since I was just using the dashboard at first. 

To set up any email forwarding rules, you'll need to create a mapping using the provided endpoints. see the rest folder for examples. You'll need to add a mapping to your kv. When an email passes through, it will check for the sender address and see if it matches to any of the mappings. Emails will only be forwarded if they are not found to be malicious.
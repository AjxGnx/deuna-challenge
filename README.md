# deuna-challenge

#To Run this challenge you need a stripe test account
at ./config/config.go you need to replaced StripeKey default value to your private test api key this is required.
StripeKey  string `required:"true" split_words:"true" default:""`
at ./static/checkout.js yo need to replaced StripePublicKey value to your public test api key this is required.
const stripe = Stripe("your public test api key");

#to run 
go run cmd/main.go

# test payments with cards
you need have the server running to see this
http://localhost:8080/api/exercise/static/checkout.html

#api docs
you need have the server running to see docs
http://localhost:8080/api/exercise/swagger/index.html



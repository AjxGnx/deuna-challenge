// This is your test publishable API key.
const stripe = Stripe("your public test api key");

const request = {
  amount : 1.25,
  customer_id:1,
  merchant_id:1
}

let elements;
let transactionID;
initialize();
checkStatus();

document
  .querySelector("#payment-form")
  .addEventListener("submit", handleSubmit);

// Fetches a payment intent and captures the client secret
async function initialize() {
  const response = await fetch("/api/exercise/transactions/stripe/intent", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify( request ),
  });
  const { data } = await response.json();


  const appearance = {
    theme: 'stripe',
  };
  const clientSecret = data.client_secret
  transactionID = data.id
  elements = stripe.elements({ appearance, clientSecret });

  const paymentElementOptions = {
    layout: "tabs",
  };

  const paymentElement = elements.create("payment", paymentElementOptions);
  paymentElement.mount("#payment-element");
}

async function handleSubmit(e) {
  e.preventDefault();
  setLoading(true);

  const { paymentIntent, error } = await stripe.confirmPayment({
    elements,
    redirect: 'if_required',
  });

  if (paymentIntent){
    await update(paymentIntent.status, paymentIntent.payment_method_types.toString(),{})
    window.location.href = "http://localhost:8080/api/exercise/static/success.html";
  }

  if (error) {
    showMessage(error.message);
    let errorJson = {
      code: error.code,
      type: error.type,
      message: error.message
    }
    await update('failed', error.payment_method.type, errorJson)
  }

  setLoading(false);
}

// Fetches the payment intent status after payment submission
async function checkStatus() {
  const clientSecret = new URLSearchParams(window.location.search).get(
    "payment_intent_client_secret"
  );

  if (!clientSecret) {
    return;
  }

  const { paymentIntent } = await stripe.retrievePaymentIntent(clientSecret);
  console.log(paymentIntent.status)
  switch (paymentIntent.status) {
    case "succeeded":
      showMessage("Payment succeeded!");
      break;
    case "processing":
      showMessage("Your payment is processing.");
      break;
    case "requires_payment_method":
      showMessage("Your payment was not successful, please try again.");
      break;
    default:
      showMessage("Something went wrong.");
      break;
  }
}

// ------- UI helpers -------

function showMessage(messageText) {
  const messageContainer = document.querySelector("#payment-message");

  messageContainer.classList.remove("hidden");
  messageContainer.textContent = messageText;

  setTimeout(function () {
    messageContainer.classList.add("hidden");
    messageContainer.textContent = "";
  }, 4000);
}

// Show a spinner on payment submission
function setLoading(isLoading) {
  if (isLoading) {
    // Disable the button and show a spinner
    document.querySelector("#submit").disabled = true;
    document.querySelector("#spinner").classList.remove("hidden");
    document.querySelector("#button-text").classList.add("hidden");
  } else {
    document.querySelector("#submit").disabled = false;
    document.querySelector("#spinner").classList.add("hidden");
    document.querySelector("#button-text").classList.remove("hidden");
  }
}

async function update(status, payment_method, error){
  let requestToUpdateTransaction = {
    status : status,
    payment_method: payment_method,
    error_code:error.code,
    error_type: error.type,
    error_message: error.message
  };


  await fetch(`/api/exercise/transactions/${transactionID}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify( requestToUpdateTransaction ),
  });
}
CLIENT_ID="$CLIENT_ID"
PROJECT_ID="$PROJECT_ID"
CLIENT_SECRET="$CLIENT_SECRET"

jq -n '{web: [ {client_id: env.CLIENT_ID, project_id: env.PROJECT_ID, auth_uri: "https://accounts.google.com/o/oauth2/auth", token_uri: "https://oauth2.googleapis.com/token", auth_provider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs", client_secret: env.CLIENT_SECRET, redirect_uris: ["http://localhost:8080/oauth/drive/redirect"]}]}' > credentials.json

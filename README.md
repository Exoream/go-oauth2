## Google OAuth2 Credential Setup
1. **Visit Google Developer Console:**
   - Go to [Google Developer Console](https://console.developers.google.com/).

2. **Create a New Project:**
   - Click on the dropdown on the top left and select "New Project."
   - Give your project a name and click "Create."

3. **Credentials:**
   - Click on the left menu, navigate to "Credentials."
   - Click "Create Credentials."
   - Choose "OAuth client ID."
   - Select the appropriate application type (e.g., "Web application").
   - **Authorized redirect URIs:**
      - Add "http://localhost:8080/callback-gl" as the authorized redirect URI.
      - If you're running your application on "http://localhost:8080," you can customize this URI to match the callback endpoint in your application.

   _Note: The authorized redirect URI is a security measure. Ensure that it matches the callback endpoint in your application. If you choose to use a different endpoint, update it accordingly._

5. **Get your Credentials:**
   - After completing the steps, click "Create" to generate your credentials.
   - Copy google client id and google client secret
   - Update your `.env` file with the google client id and google client secret
  
```dotenv
GOOGLE_CLIENT_ID=your-client-id
GOOGLE_CLIENT_SECRET=your-client-secret
CALLBACK_URL=http://localhost:8080/callback-gl

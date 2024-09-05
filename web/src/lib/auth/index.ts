import NextAuth from "next-auth";
import keycloak from "next-auth/providers/keycloak";

export const {
  handlers,
  auth,
  signIn,
  signOut,
} = NextAuth({
  providers: [keycloak({
    clientId: process.env.KEYCLOAK_ID,
    clientSecret: process.env.KEYCLOAK_SECRET,
    issuer: process.env.KEYCLOAK_ISSUER,
  })],
  callbacks: {
    jwt: async ({ token, user, account, profile }) => {
      if (account && user) {
        token.accessToken = account.access_token
        token.idToken = account.id_token
        token.refreshToken = account.refresh_token
        token.tokenType = account.token_type
        token.user = user
      }
      return token
    },
    redirect: async ({ url, baseUrl }) => {
      console.log('url:', url);
      console.log('baseUrl:', baseUrl);
      return Promise.resolve(url)
    },
  }
});
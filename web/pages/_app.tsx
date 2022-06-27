import { ProvideAuth } from "@/components/auth/AuthRoute";
import { AppProps } from "next/app";
import Head from "next/head";

// import "tailwindcss/tailwind.css";
import "../styles/globals.css";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ProvideAuth>
      <Head>
        <title>Workflow</title>
        <link rel="icon" href="/favicon.ico" />
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <meta
          name="description"
          content="Web site created using create-react-app"
        />
      </Head>

      <Component {...pageProps} />
    </ProvideAuth>
  );
}

export default MyApp;

import { Head } from "$fresh/runtime.ts";
import { useSignal } from "@preact/signals";

export default function Home() {
  const count = useSignal(3);
  return (
    <>
      <Head>
        <title>E - commerce</title>
      </Head>
      <div class="p-4 mx-auto max-w-screen-md">
        <h1>Hello Deno Fresh</h1>
        <h3>{count}</h3>
      </div>
    </>
  );
}

import { generateKey, readKey, readPrivateKey } from "openpgp";
import { KeyPairType } from "./types";

export async function initKeys(): Promise<KeyPairType> {
  const rawPrivateKey = localStorage.getItem("neutrino-private-key");
  const rawPublicKey = localStorage.getItem("neutrino-public-key");

  if (rawPrivateKey !== null && rawPublicKey !== null) {
    try {
      const privateKey = await readPrivateKey({ armoredKey: rawPrivateKey });
      const publicKey = await readKey({ armoredKey: rawPublicKey });

      const exp = await publicKey.getExpirationTime();
      const now = new Date();
      if (exp !== null && exp > now) {
        // The keypair is still valid, we can use it
        console.log(
          "A key pair was found in the local storage, and is still valid"
        );
        return { publicKey, privateKey };
      }
    } catch (e) {
      console.warn(
        "A key pair was found in the local storage but is not valid. It has been discarded."
      );
    }
  }

  try {
    // Generate a private key
    const keyPair = await generateKey({
      type: "ecc",
      curve: "p256",
      keyExpirationTime: 3600 * 24 * 30 + 3600, // Default duration of a server session, plus a buffer to be safe
      userIDs: [{ name: "anonymous" }],
      format: "armored",
      // TODO: maybe handle the subkey generation
    });
    const publicKey = await readKey({ armoredKey: keyPair.publicKey });
    const privateKey = await readPrivateKey({ armoredKey: keyPair.privateKey });

    localStorage.setItem("neutrino-private-key", keyPair.privateKey);
    localStorage.setItem("neutrino-public-key", keyPair.publicKey);

    return { publicKey, privateKey };
  } catch (e) {
    throw new Error("The PGP key pair could not be generated: " + String(e));
  }
}

import { defineStore } from 'pinia'
import { generateKey, readKey, readPrivateKey } from "openpgp";
import { KeyPairType } from "@/types";

function getCookie(cookieName: string): string | undefined {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${cookieName}=`);
  if (parts.length === 2) {
    const cookie = parts.pop();
    if (cookie !== undefined) {
      return cookie.split(';').shift();
    }
  }
}

export const useNeutrinoStore = defineStore('neutrino', {
  state: () => ({
    nickname: null as string | null,
    keyPair: {
      publicKey: null,
      privateKey: null,
    } as KeyPairType,
    initError: null as string | null,
    initHasRun: false,
  }),

  actions: {
    async initializeData() {
      // The method can only run once
      if (this.initHasRun) {
        return
      }
      this.initHasRun = true;

      const sessionCookie = getCookie("neutrino-js-cookie");
      if (sessionCookie !== undefined) {
        // There is a session, we extract the keypair from local storage
        const rawPrivateKey = localStorage.getItem("neutrino-private-key");
        const rawPublicKey = localStorage.getItem("neutrino-public-key");
        const nickname = localStorage.getItem("neutrino-nickname");

        if (nickname != null && rawPrivateKey !== null && rawPublicKey !== null) {
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
              this.nickname = nickname;
              this.keyPair = { publicKey, privateKey };
              return;
            }
          } catch (e) {
            console.warn("A key pair was found in the local storage but is not valid. It has been discarded.");
          }
        }
      }

      localStorage.clear();

      try {
        // Generate a private key
        const keyPair = await generateKey({
          type: "ecc",
          curve: "p256",
          keyExpirationTime: 3600 * 24 * 30 + 3600, // Default duration of a server session, plus a buffer to be safe
          userIDs: [{ name: "anonymous" }],
          format: "armored",
        });
        const publicKey = await readKey({ armoredKey: keyPair.publicKey });
        const privateKey = await readPrivateKey({ armoredKey: keyPair.privateKey });

        localStorage.setItem("neutrino-private-key", keyPair.privateKey);
        localStorage.setItem("neutrino-public-key", keyPair.publicKey);

        this.nickname = null;
        this.keyPair = { publicKey, privateKey };
        return
      } catch (e) {
        const errMsg = "The PGP key pair could not be generated: " + String(e)
        console.warn("Initialization error: " + errMsg)
        this.initError = errMsg
      }
    },

    setNickname(newNickname: string) {
      this.nickname = newNickname;
      localStorage.setItem("neutrino-nickname", newNickname);
    },
  }
})

import { Button, Label, Modal, TextInput } from "flowbite-react";
import { useRef } from "react";
import { auth } from "../util/firebase";
import { signInWithEmailAndPassword } from "firebase/auth";


export default function SigninModal({ openModal, setOpenModal, onSignInComplete }: { openModal: boolean, setOpenModal: (open: boolean) => void, onSignInComplete: (message: string) => void}) {
  const emailInputRef = useRef<HTMLInputElement>(null);
  const passwordInputRef = useRef<HTMLInputElement>(null);
  return (
        <Modal show={openModal} size="md" popup onClose={() => setOpenModal(false)} initialFocus={emailInputRef}>
        <Modal.Header />
        <Modal.Body>
          <div className="space-y-6">
            <h3 className="text-xl font-medium text-gray-900 dark:text-white">Sign in to your account</h3>
            <div>
              <div className="mb-2 block">
                <Label htmlFor="email" value="Your email" />
              </div>
              <TextInput id="email" ref={emailInputRef} placeholder="name@company.com" required />
            </div>
            <div>
              <div className="mb-2 block">
                <Label htmlFor="password" value="Your password" />
              </div>
              <TextInput id="password" ref={passwordInputRef} type="password" required />
            </div>
            
            <div className="w-full">
              <Button className="bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange" onClick={() => {
                
                signInWithEmailAndPassword(auth, emailInputRef.current?.value as string, passwordInputRef.current?.value as string)
                .then((userCredential) => {
                  const user = userCredential.user;
                  setOpenModal(false)
                  onSignInComplete(`Signed in as ${user?.email}`)
                }).catch((error) => {
                  const errorCode = error.code;
                  const errorMessage = error.message;
                  onSignInComplete(`Error: ${errorCode} ${errorMessage}`)
                })
              }}>Log in to your account</Button>
            </div>
          </div>
        </Modal.Body>
      </Modal>
    )
}
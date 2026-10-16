import { Button, Label, Modal, ModalBody, ModalHeader, TextInput } from "flowbite-react";
import { useRef } from "react";
import { useApplication } from "@/interfaces/presentation/application-context";


export default function SigninModal({ openModal, setOpenModal, onSignInComplete }: { openModal: boolean, setOpenModal: (open: boolean) => void, onSignInComplete: (message: string) => void}) {
  const { auth } = useApplication();
  const emailInputRef = useRef<HTMLInputElement>(null);
  const passwordInputRef = useRef<HTMLInputElement>(null);
  return (
        <Modal show={openModal} size="md" popup onClose={() => setOpenModal(false)} initialFocus={emailInputRef}>
        <ModalHeader />
        <ModalBody>
          <div className="space-y-6">
            <h3 className="text-xl font-medium text-gray-900 dark:text-white">Sign in to your account</h3>
            <div>
              <div className="mb-2 block">
                <Label htmlFor="email">Your email</Label>
              </div>
              <TextInput id="email" ref={emailInputRef} placeholder="name@company.com" required />
            </div>
            <div>
              <div className="mb-2 block">
                <Label htmlFor="password">Your password</Label>
              </div>
              <TextInput id="password" ref={passwordInputRef} type="password" required />
            </div>
            
            <div className="w-full">
              <Button className="bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange" onClick={() => {
                
                auth.signIn(emailInputRef.current?.value ?? "", passwordInputRef.current?.value ?? "")
                .then((user) => {
                  setOpenModal(false)
                  onSignInComplete(`Signed in as ${user.email}`)
                }).catch((error) => {
                  const errorCode = error.code;
                  const errorMessage = error.message;
                  onSignInComplete(`Error: ${errorCode} ${errorMessage}`)
                })
              }}>Log in to your account</Button>
            </div>
          </div>
        </ModalBody>
      </Modal>
    )
}
import { Navbar, Toast } from 'flowbite-react';
import { useLocation } from 'react-router-dom';
import {MESSAGES, NOFICATIONS, PAYMENTS} from "../util/constants";
import { auth } from '../util/firebase';
import { useEffect, useState } from 'react';
import SigninModal from './signinmodal';
import MomentumFooter from './footer';
import { HiFire } from "react-icons/hi";


function AuthNavlink({ onSigninClick }: { onSigninClick: () => void}){
  const [isSignedIn, setSignedIn] = useState(auth.currentUser ? true : false)
  useEffect(() => {
    auth.onAuthStateChanged((user) => {
      if (user) {
        setSignedIn(true)
      } else {
        setSignedIn(false)
      }
    })
    if (auth.currentUser) {
      setSignedIn(true)
    } else {
      setSignedIn(false)
    }
  }, [auth.currentUser])
  if (isSignedIn) {
    return <Navbar.Link className="cursor-pointer" onClick={() => {
      auth.signOut().then(() => {
        console.log("Signed out")
      }).catch((error) => {
        alert(error)
      })
    }}>Sign Out</Navbar.Link>
  } else {
    return <Navbar.Link onClick={onSigninClick}>Sign In</Navbar.Link>
  }
}

export function MomentumNavigation({ content, showAuth}: { content?: React.ReactNode, showAuth?: boolean}) {
    const location = useLocation();
    const [openAuthModal, setOpenAuthModal] = useState(false);
    const [showToast, setShowToast] = useState(false);
    const [toastMessage, setToastMessage] = useState("");
    useEffect(() => {
        if (showAuth) {
          setOpenAuthModal(true)
        }
    }, [showAuth])
    return (
      <>
        <Navbar fluid rounded className="w-full" >
          <Navbar.Brand href="/dashboard">
            <img src="assets/logo.png" className="mr-3 h-6 sm:h-9" alt="Momentum Logo" />
            <span className="self-center whitespace-nowrap text-xl font-bold border-b border-gray-100 text-gray-700 hover:bg-gray-50 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white md:border-0 md:hover:bg-transparent md:hover:text-cyan-700 md:dark:hover:bg-transparent md:dark:hover:text-white">Momentum</span>
          </Navbar.Brand>
          <Navbar.Toggle />
          <Navbar.Collapse>
            {/* <Navbar.Link  href={HOME} active={HOME === location.pathname}> Home</Navbar.Link> */}
            <Navbar.Link href={MESSAGES} active={MESSAGES === location.pathname}>Messages</Navbar.Link>
            <Navbar.Link href={NOFICATIONS} active={NOFICATIONS === location.pathname}>Notifications</Navbar.Link>
            <Navbar.Link href={PAYMENTS} active={PAYMENTS === location.pathname}>Transactions</Navbar.Link>
            {/* <Navbar.Link href={USERS} active={USERS === location.pathname}>Users</Navbar.Link> */}
            <AuthNavlink onSigninClick={() => setOpenAuthModal(true)}/>
          </Navbar.Collapse>
        </Navbar>
        {showToast && (
          <div className='w-full flex items-center justify-end'>
            <Toast>
              <div className="inline-flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-cyan-100 text-cyan-500 dark:bg-cyan-800 dark:text-cyan-200">
                <HiFire className="h-5 w-5" />
              </div>
              <div className="ml-3 text-sm font-normal">{toastMessage}</div>
              <Toast.Toggle onDismiss={() => setShowToast(false)} />
            </Toast>
          </div>
        )}
        <div className="w-full h-screen flex flex-col justify-between items-center">
          <div> {content} </div>
          <MomentumFooter />
        </div>
        <SigninModal 
          openModal={openAuthModal} 
          setOpenModal={setOpenAuthModal} 
          onSignInComplete={(message) => {
            setToastMessage(message)
            setShowToast(true)
          }}
        />
      </>
    );
  }
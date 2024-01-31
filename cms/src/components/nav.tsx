import { Navbar } from 'flowbite-react';
import { useLocation } from 'react-router-dom';
import {MESSAGES, NOFICATIONS, PAYMENTS} from "../util/constants";
import { auth } from '../util/firebase';
import { useEffect, useState } from 'react';


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
        alert("Signed out")
      }).catch((error) => {
        alert(error)
      })
    }}>Sign Out</Navbar.Link>
  } else {
    return <Navbar.Link onClick={onSigninClick}>Sign In</Navbar.Link>
  }
}

export function Nav({ onSigninClick }: { onSigninClick: () => void }) {
  let location = useLocation();
  
    return (
      <Navbar fluid rounded className="w-full" >
        <Navbar.Brand href="/dashboard">
          <img src="assets/logo.png" className="mr-3 h-6 sm:h-9" alt="Momentum Logo" />
          <span className="self-center whitespace-nowrap text-xl font-bold border-b border-gray-100 text-gray-700 hover:bg-gray-50 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white md:border-0 md:hover:bg-transparent md:hover:text-cyan-700 md:dark:hover:bg-transparent md:dark:hover:text-white">Momentum Dashboard</span>
        </Navbar.Brand>
        <Navbar.Toggle />
        <Navbar.Collapse>
          {/* <Navbar.Link  href={HOME} active={HOME === location.pathname}> Home</Navbar.Link> */}
          <Navbar.Link href={MESSAGES} active={MESSAGES === location.pathname}>Messages</Navbar.Link>
          <Navbar.Link href={NOFICATIONS} active={NOFICATIONS === location.pathname}>Notifications</Navbar.Link>
          <Navbar.Link href={PAYMENTS} active={PAYMENTS === location.pathname}>Transactions</Navbar.Link>
          {/* <Navbar.Link href={USERS} active={USERS === location.pathname}>Users</Navbar.Link> */}
          <AuthNavlink onSigninClick={onSigninClick}/>
        </Navbar.Collapse>
      </Navbar>
    );
  }
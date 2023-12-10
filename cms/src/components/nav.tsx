import { Navbar } from 'flowbite-react';
import { useLocation } from 'react-router-dom';
import {MESSAGES, NOFICATIONS, PAYMENTS, USERS} from "../util/constants";




export function Nav() {
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
          <Navbar.Link  className="hover:text-momentum-orange active:text-momentum-orange" href={MESSAGES} active={MESSAGES === location.pathname}>Messages</Navbar.Link>
          <Navbar.Link href={NOFICATIONS} active={NOFICATIONS === location.pathname}>Notifications</Navbar.Link>
          <Navbar.Link href={PAYMENTS} active={PAYMENTS === location.pathname}>Payments</Navbar.Link>
          <Navbar.Link href={USERS} active={USERS === location.pathname}>Users</Navbar.Link>
        </Navbar.Collapse>
      </Navbar>
    );
  }
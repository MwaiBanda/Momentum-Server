import { Navbar } from 'flowbite-react';

export function Nav() {
    return (
      <Navbar fluid rounded className="w-full">
        <Navbar.Brand href="/dashboard">
          <img src="assets/logo.png" className="mr-3 h-6 sm:h-9" alt="Momentum Logo" />
          <span className="self-center whitespace-nowrap text-xl font-bold border-b border-gray-100 text-gray-700 hover:bg-gray-50 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white md:border-0 md:hover:bg-transparent md:hover:text-cyan-700 md:dark:hover:bg-transparent md:dark:hover:text-white">Momentum Dashboard</span>
        </Navbar.Brand>
        <Navbar.Toggle />
        <Navbar.Collapse>
          <Navbar.Link href="/dashboard" active> Home</Navbar.Link>
          <Navbar.Link href="/messages">Messages</Navbar.Link>
          <Navbar.Link href="/notifications">Notifications</Navbar.Link>
          <Navbar.Link href="/payments">Payments</Navbar.Link>
          <Navbar.Link href="/users">Users</Navbar.Link>
        </Navbar.Collapse>
      </Navbar>
    );
  }
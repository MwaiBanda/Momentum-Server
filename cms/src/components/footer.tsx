import { Footer } from "flowbite-react";

export default function MomentumFooter() {
  return (
    <Footer container className="mt-5 mb-5">
      <div className="w-full text-center">
        <div className="w-full justify-between sm:flex sm:items-center sm:justify-between">
          <Footer.Brand
            href="/dashboard"
            src="assets/logo.png"
            alt="Momentum Logo"
            name="Momentum"
          />
          <Footer.LinkGroup>
            <Footer.Link href="/messages">Messages</Footer.Link>
            <Footer.Link href="/notifications">Notifications</Footer.Link>
            <Footer.Link href="/payments">Transactions</Footer.Link>
          </Footer.LinkGroup>
        </div>
        <Footer.Divider />
        <Footer.Copyright href="/dashboard" by="Momentum Church™" year={2024} />
      </div>
    </Footer>
  );
}
import {
  Footer,
  FooterBrand,
  FooterCopyright,
  FooterDivider,
  FooterLink,
  FooterLinkGroup,
} from "flowbite-react";

export default function MomentumFooter() {
  return (
    <Footer container className="mt-5 mb-5">
      <div className="w-full text-center">
        <div className="w-full justify-between sm:flex sm:items-center sm:justify-between">
          <FooterBrand
            href="/dashboard"
            src="assets/logo.png"
            alt="Momentum Logo"
            name="Momentum"
          />
          <FooterLinkGroup>
            <FooterLink href="/messages">Messages</FooterLink>
            <FooterLink href="/notifications">Notifications</FooterLink>
            <FooterLink href="/payments">Transactions</FooterLink>
          </FooterLinkGroup>
        </div>
        <FooterDivider />
        <FooterCopyright href="/dashboard" by="Momentum Church™" year={2024} />
      </div>
    </Footer>
  );
}
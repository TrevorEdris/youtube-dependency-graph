import logo from '../logo.svg';
import { Container, Nav, Navbar, NavDropdown } from 'react-bootstrap';

function YDGNavbar() {
	return (
		<Navbar bg="dark" variant="dark" expand="lg">
			<Container>
				<Navbar.Brand href="#home">
					{' '}
					<img alt="" src={logo} width="30" height="30" className="d-inline-block align-top" />YDG
				</Navbar.Brand>
				<Navbar.Toggle aria-controls="basic-navbar-nav" />
				<Navbar.Collapse id="basic-navbar-nav">
					<Nav className="me-auto">
						<NavDropdown title="New" id="basic-nav-dropdown">
							<NavDropdown.Item href="#from_title">From Title</NavDropdown.Item>
							<NavDropdown.Item href="#from_url">From URL</NavDropdown.Item>
							<NavDropdown.Item href="#from_id">From ID</NavDropdown.Item>
						</NavDropdown>
						<Nav.Link href="#import">Import</Nav.Link>
					</Nav>
				</Navbar.Collapse>
			</Container>
		</Navbar>
	);
}

export default YDGNavbar;

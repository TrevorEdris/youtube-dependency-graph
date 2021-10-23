import logo from '../logo.svg';
import { Container, Nav, Navbar, NavDropdown } from 'react-bootstrap';

function fromTitle(event) {
	event.preventDefault();
	console.log('fromTitle');
	console.log(event);
}

function fromURL(event) {
	event.preventDefault();
	console.log('fromURL');
	console.log(event);
}

function fromID(event) {
	event.preventDefault();
	console.log('fromID');
	console.log(event);
}

function importGraph(event) {
	event.preventDefault();
	console.log('importGraph');
	console.log(event);
}

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
							<NavDropdown.Item href="#from_title" onClick={fromTitle}>
								From Title
							</NavDropdown.Item>
							<NavDropdown.Item href="#from_url" onClick={fromURL}>
								From URL
							</NavDropdown.Item>
							<NavDropdown.Item href="#from_id" onClick={fromID}>
								From ID
							</NavDropdown.Item>
						</NavDropdown>
						<Nav.Link href="#import" onClick={importGraph}>
							Import
						</Nav.Link>
					</Nav>
				</Navbar.Collapse>
			</Container>
		</Navbar>
	);
}

export default YDGNavbar;

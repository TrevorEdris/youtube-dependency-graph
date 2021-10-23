import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';

import YDGNavbar from './components/navbar';
import GraphPane from './components/graph_pane';

// https://www.codeply.com/go/Iyjsd8djnz
function App() {
	return (
		<div className="App container-fluid vh-100">
			<div className="row justify-content-center h-100">
				<div className="col">
					<div className="h-100 d-flex flex-column">
						<div className="row">
							<YDGNavbar />
						</div>
						<div className="row border border-dark flex-grow-1">
							<GraphPane />
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}

export default App;

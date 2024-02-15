import Home from "./Pages/Home.js"
import CurrentSignoutsDisplay from "./Components/CurrentSignoutsDisplay.js";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import HistoryPage from "./Pages/History.js";
import AboutPage from "./Pages/About.js";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/home" element={<Home />}>
            <Route index element={<CurrentSignoutsDisplay />} />
            <Route path="/home/about" element={<AboutPage />} />
            <Route path="/home/history" element={<HistoryPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;

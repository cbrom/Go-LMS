import React from 'react';

import Navbar from './NavBar';
import Section1 from './Section1';
import Section2 from './Section2';
import Section3 from './Section3';
import Section4 from './Section4';
import Footer from './Footer';


const App = () => {
    return (
        <div>
            <Navbar />
            <Section1 />
            <Section2 />
            <Section3 />
            <Section4 />
            <Footer />
        </div>
    );
}

export default App;
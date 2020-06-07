// import './NavBar.css';
import React from 'react';
import logo from '../img/logo.svg';
import './NavBar2.css';

const Navbar = () => {
    return (
        <div id="pupilfirst-header">
            <nav className="bg-white border-b border-gray-200 max-w-6xl w-full mx-auto px-3 md:px-4 mt-2"><div className="max-w-7xl mx-auto">
                <div className="flex justify-between h-16 w-full">
                    <div className="flex w-full justify-between">
                        <a className="flex-shrink-0 flex items-center" href="./">
                            <img className="w-26 md:w-38" alt="Pupilfirst Logo" src={logo} />
                        </a>
                        <div className="hidden sm:-my-px sm:ml-6 sm:flex">
                            <a className="ml-8 inline-flex items-center px-1 pt-1 text-sm font-semibold leading-5 transition duration-150 ease-in-out border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 focus:outline-none focus:text-gray-700 focus:border-gray-300 " href="/pricing">Pricing</a>
                            <a className="ml-8 inline-flex items-center px-1 pt-1 text-sm font-semibold leading-5 transition duration-150 ease-in-out border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 focus:outline-none focus:text-gray-700 focus:border-gray-300 " href="/support">Support</a>
                            <a className="ml-8 inline-flex items-center px-1 pt-1 text-sm font-semibold leading-5 transition duration-150 ease-in-out border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 focus:outline-none focus:text-gray-700 focus:border-gray-300 " href="https://blog.pupilfirst.com/">Blog</a>
                            <a className="ml-8 inline-flex items-center px-1 pt-1 text-sm font-semibold leading-5 transition duration-150 ease-in-out border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 focus:outline-none focus:text-gray-700 focus:border-gray-300 " href="https://docs.pupilfirst.com">Docs</a>
                            <a className="ml-8 inline-flex items-center px-1 pt-1 text-sm font-semibold leading-5 transition duration-150 ease-in-out border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 focus:outline-none focus:text-gray-700 focus:border-gray-300 " href="/sign_in">Sign in</a>
                            <div className="inline-flex items-center">
                                <a className="btn btn-success ml-2 md:ml-4" href="/sign_up">Start Free Trial</a>
                            </div>
                        </div>
                    </div>
                    <div className="-mr-2 flex items-center sm:hidden">
                        <a className="btn btn-success btn-small" href="/sign_up">Start Free Trial</a>
                            <button className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 focus:text-gray-500 transition duration-150 ease-in-out">
                                <svg className="block h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path d="M4 6h16M4 12h16M4 18h16" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2">
                                    </path>
                                </svg>
                            </button>
                    </div>
                </div>
                </div>
            </nav>
        </div>
    );
}

export default Navbar;
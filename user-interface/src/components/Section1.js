import React from 'react';
import './NavBar2.css';
import './NavBar.css';
import pupilFirstHeroImage from '../pupilfirst-hero-illustration.svg';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBolt } from '@fortawesome/free-solid-svg-icons';
import { faGithub } from '@fortawesome/free-brands-svg-icons';

const Section1 = () => {
    return (
        <section className="pf-landing-hero border-b flex">
            <div className="container max-w-6xl mx-auto px-4 py-5 md:py-0 border-t md:border-0 flex items-center flex-wrap overflow-hidden text-center md:text-left">
            <div className="w-full lg:w-1/2">
                <span className="pf-landing__hero-label pf-landing__hero-label--highlight relative z-0 md:text-xl md:ml-1">
                The open-source LMS that lets you
                </span>
                <h1 className="pf-landing__hero-heading font-bold">Focus on your
                <span className="block pf-landing__hero-heading-highlight font-extrabold text-primary-500">students.</span></h1>

                <p className="leading-normal mt-4 md:pr-16">Teaching is difficult. Teaching online even more so.
                Pupilfirst gives you a proven method to keep your students engaged while keeping you, the teacher, always in
                the loop.</p>

                <div className="flex flex-wrap justify-center md:justify-start mt-6 md:mt-10">
                    <a href="/sign_up" className="btn btn-primary btn-large w-full md:w-auto md:mr-4">
                    <FontAwesomeIcon icon={faBolt} size="lg" fixedWidth />
                    <span className="ml-2">Try our hosted solution for FREE</span>
                    </a>
                <a href="https://github.com/pupilfirst/pupilfirst" className="btn btn-primary-ghost btn-large w-full md:w-auto mt-4 sm:mt-0">
                    <FontAwesomeIcon icon={faGithub} size="lg" fixedWidth />
                    <span className="ml-2">Deploy on your server</span>
                </a>
                </div>
            </div>

            <div className="w-full lg:w-1/2">
                <img className="w-full" alt="Pupilfirst hero" src={pupilFirstHeroImage} />
            </div>
            </div>
        </section>
    );
}

export default Section1;
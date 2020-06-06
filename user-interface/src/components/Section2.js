import React from 'react';

import tasksFeatureIcon from '../img/tasks-feature-icon.svg';
import feedbackFeatureIcon from '../img/feedback-feature-icon.svg';
import communityFeatureIcon from '../img/community-feature-icon.svg';
import whitelabelFeatureIcon from '../img/whitelabel-feature-icon.svg';

const Section2 = () => {
    return (
        <section className="bg-gray-100 flex px-3 pt-8 pb-10">
        <div className="max-w-4xl mx-auto text-center">
        <h2 className="pf-landing__section-heading inline-block relative z-0 pf-landing__section-heading--highlight text-2xl md:text-4xl text-center">
            What You Can Do with Pupilfirst
        </h2>
        <div className="flex flex-wrap text-center mt-6 md:mt-8">
            <div className="w-full sm:w-1/2">
            <div className="pf-landing__feature-card h-full bg-white md:bg-gray-100 rounded-lg md:rounded-none border md:border-l-0 md:border-t-0 md:border-b-0 md:border-r border-gray-400 p-6 md:py-8 md:px-12">
                <img className="w-24 mx-auto" alt="Present actionable content" src={tasksFeatureIcon} />
                <h4 className="text-lg leading-snug mt-3">Present actionable <span className="block">content</span></h4>
                <p className="mt-2 text-sm">Pupilfirst is built around the philosophy that true learning cannot happen by just
                consuming information. It happens by students attempting relevant tasks and learning the theory around the
                tasks.</p>
            </div>
            </div>
            <div className="w-full sm:w-1/2 mt-4 md:mt-0">
            <div className="pf-landing__feature-card h-ful bg-white md:bg-gray-100 rounded-lg md:rounded-none border border-gray-400 md:border-0 p-6 md:py-8 md:px-12">
                <img className="w-24 mx-auto" alt="Give targeted feedback" src={feedbackFeatureIcon} />
                <h4 className="text-lg leading-snug mt-3">Give targeted <span className="block">feedback</span></h4>
                <p className="mt-2 text-sm">You are an expert in your domain. Pupilfirst enables you to quickly and efficiently
                review the quality of task submissions by students, share feedback and create a conversation around what
                students have learned.</p>
            </div>
            </div>
            <div className="flex flex-wrap md:border-t border-gray-400">
            <div className="w-full sm:w-1/2 mt-4 md:mt-0">
                <div className="pf-landing__feature-card h-full bg-white md:bg-gray-100 rounded-lg md:rounded-none border md:border-l-0 md:border-t-0 md:border-b-0 md:border-r border-gray-400 p-6 md:py-8 md:px-10">
                <img className="w-24 mx-auto" alt="Foster a culture of sharing knowledge" src={communityFeatureIcon} />
                <h4 className="text-lg leading-snug mt-3">Foster a culture of sharing <span className="block">knowledge</span>
                </h4>
                <p className="mt-2 text-sm">Pupilfirst encourages collaboration by building student communities that can solve
                    its own problems, and even supports teaming up of students to go through a challenging course
                    together.</p>
                </div>
            </div>
            <div className="w-full sm:w-1/2 mt-4 md:mt-0">
                <div className="pf-landing__feature-card h-full bg-white md:bg-gray-100 rounded-lg md:rounded-none border border-gray-400 md:border-0 p-6 md:py-8 md:px-12">
                <img className="w-24 mx-auto" alt="Build your own brand" src={whitelabelFeatureIcon} />
                <h4 className="text-lg leading-snug mt-3">Build your own <span className="block">brand</span></h4>
                <p className="mt-2 text-sm">Use your own identity when deploying Pupilfirst - customize what your student sees
                    to leverage your brand for greater reach and build a trusted online presence.</p>
                </div>
            </div>
            </div>
        </div>
        </div>
    </section>
    );
}

export default Section2;
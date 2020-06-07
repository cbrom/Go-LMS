import React from 'react';

const Section3 = () => {
    return (
    <section className="bg-white border-t border-b flex flex-col px-3 py-8 md:pb-10">
        <div className="max-w-3xl mx-auto w-full text-center">
        <h2 className="pf-landing__section-heading inline-block relative z-0 pf-landing__section-heading--highlight text-2xl md:text-4xl text-center">Watch
            Pupilfirst in Action</h2>
        <div className="pf-landing__video-embed rounded-lg mt-6">
            <iframe width="560" height="315" src="https://www.youtube.com/embed/il0L60hfVEk" frameBorder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowFullScreen=""></iframe>
        </div>
        </div>
    </section>
    );
}

export default Section3;
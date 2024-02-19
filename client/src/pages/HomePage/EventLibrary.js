import React from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import { Autoplay, Navigation } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/navigation';
import { useInView } from 'react-intersection-observer';
import { motion, AnimatePresence } from 'framer-motion';

import event1 from "./images/next-event-1.jpg";
import event2 from "./images/next-event-2.jpg";
import event3 from "./images/next-event-3.jpg";

function HomepageNextEvents() {
    const events = [
        { img: event1, title: "Welcoming Party 2018", description: "Green Palace, 22 Street, 23-28, Los Angeles California" },
        { img: event2, title: "Welcoming Party 2018", description: "Green Palace, 22 Street, 23-28, Los Angeles California" },
        { img: event3, title: "Welcoming Party 2018", description: "Green Palace, 22 Street, 23-28, Los Angeles California" },
        { img: event1, title: "Welcoming Party 2018", description: "Green Palace, 22 Street, 23-28, Los Angeles California" },
        { img: event2, title: "Welcoming Party 2018", description: "Green Palace, 22 Street, 23-28, Los Angeles California" },
        { img: event3, title: "Welcoming Party 2018", description: "Green Palace, 22 Street, 23-28, Los Angeles California" },
        // Thêm các sự kiện khác tương tự
    ];

    const { ref, inView } = useInView({
        triggerOnce: true,
        threshold: 0.5,
    });

    return (
        <div ref={ref}>
            <AnimatePresence>
                {inView && (
                    <motion.div
                        initial={{ opacity: 0, translateY: 200 }}
                        animate={{ opacity: 1, translateY: 0 }}
                        exit={{ opacity: 0, translateY: -50 }}
                        transition={{ duration: 1 }}
                    >
                        <div className="container mb-5">
                            <div className="row">
                                <div className="col-12">
                                    <div className="entry-title">
                                        <p>Job Fair</p>
                                        <h2>Hình ảnh sự kiện</h2>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <Swiper
                            modules={[Autoplay, Navigation]}
                            spaceBetween={50}
                            slidesPerView={3}
                            loop={true}
                            centeredSlides={true}
                            navigation
                            autoplay={{
                                delay: 3000,
                                disableOnInteraction: false,
                            }}
                            className="mySwiper"
                        >
                            {events.map((event, index) => (
                                <SwiperSlide key={index}>
                                    <div className="next-event-content">
                                        <figure className="featured-image">
                                            <img src={event.img} alt={event.title} />
                                            <div className="entry-content flex flex-column justify-content-center align-items-center">
                                                {/* <h3>{event.title}</h3> */}
                                                {/* <p>{event.description}</p> */}
                                            </div>
                                        </figure>
                                    </div>
                                </SwiperSlide>
                            ))}
                        </Swiper>
                    </motion.div>
                )}
            </AnimatePresence>
        </div>
    );
}

export default HomepageNextEvents;

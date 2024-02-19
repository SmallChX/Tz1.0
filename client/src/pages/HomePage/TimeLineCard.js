// ArtistCard.js
import React from 'react';
import { useInView } from 'react-intersection-observer';
import { motion } from 'framer-motion';
import { Button } from 'react-bootstrap';

const ArtistCard = ({ artist, index }) => {
    const { ref, inView } = useInView({
        triggerOnce: true, // Chỉ kích hoạt một lần
        threshold: 0.2,  // Phần tử xuất hiện 30% mới bắt đầu hiệu ứng
    });

    return (
        <motion.div ref={ref}
                    initial={{ opacity: 0, x: index % 2 === 0 ? -100 : 100 }}
                    animate={inView ? { opacity: 1, x: 0 } : {}}
                    transition={{ duration: 0.5, delay: index * 0.2 }}
                    className="mb-4 lineup-artists-wrap flex flex-wrap">
            {index % 2 === 0 ? (
                        <>
                        <img className="featured-image"variant="top " src={artist.image} alt={artist.name}/>
                        <div className="lineup-artists-description">
                            <div className='lineup-artists-description-container'>
                                <div className="entry-title">{artist.name}</div>
                                <div className="entry-content">{artist.description}</div>
                                {/* <Button variant="primary" href={artist.link}>Learn More</Button> */}
                            </div>
                        </div>
                        </>
                    ) : (
                        <>
                        <div className="lineup-artists-description">
                            <div className='lineup-artists-description-container'>
                                <div className="entry-title">{artist.name}</div>
                                <div className="entry-content">{artist.description}</div>
                                {/* <Button variant="primary" href={artist.link}>Learn More</Button> */}
                            </div>
                        </div>
                        <img className="featured-image"variant="top " src={artist.image} alt={artist.name} />
                        </>
                    )}
        </motion.div>
    );
}

export default ArtistCard;

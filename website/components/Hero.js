import Head from 'next/head';

const Hero = () => {
  return (
    <div className='hero'>
      <Head>
        <title>Cric | Get live cricket score in your terminal</title>
      </Head>
      <h1>🏏💻 Cric</h1>
      <h4 className='subheading'>
        For all the cricket lovers who happen to be on the terminal all the
        time.
      </h4>
    </div>
  );
};

export default Hero;

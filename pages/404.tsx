import type { NextPage } from 'next'
import Head from 'next/head'
import styles from '../styles/Home.module.css'

const Home: NextPage = () => {
	return (
		<div className={styles.container}>
			<Head>
				<title>Toby Scott - Not Found</title>
				<meta name="description" content="The page you requested doesn&apos;t exist" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<link rel="icon" href="/favicon.ico" />
			</Head>

			<div className={styles.main}>

				<div className="my-5 text-center">
					<div className="text-6xl">404 - Not Found</div>
					<div className="text-2xl mt-4">The page you requested doesn&apos;t exist</div>
				</div>
				
			</div>

		</div>
	)
}

export default Home
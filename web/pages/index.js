import React, { useState, useEffect } from "react";

import Head from "next/head";

import { Dashboard } from "../components";

const { API_ENDPOINT } = process.env

const IndexPage = () => {
  const [numberOfDailyMessage, setNumberOfDailyMessage] = useState(0)
  const [topMessagesByEngagements, setTopMessagesByEngagements] = useState([])
  const [topAccountsByMessages, setTopAccountsByMessages] = useState([])
  const [words, setWords] = useState([])
  const [hashtags, setHashtags] = useState([])

  useEffect(() => {
    fetch(`${API_ENDPOINT}/number_of_daily_message`)
      .then(res => res.json())
      .then(json => {
        setNumberOfDailyMessage(json["count"]);
      })
      .catch(console.error);
    fetch(`${API_ENDPOINT}/top_messages_by_engagements`)
      .then(res => res.json())
      .then(json => {
        setTopMessagesByEngagements(json["top_messages_by_engagements"]);
      })
      .catch(console.error);
    fetch(`${API_ENDPOINT}/top_accounts_by_messages`)
      .then(res => res.json())
      .then((json) => {
        setTopAccountsByMessages(json['top_accounts_by_messages'])
      })
      .catch(console.error)
    fetch(`${API_ENDPOINT}/words`)
      .then(res => res.json())
      .then(({ words }) => setWords(words))
      .catch(console.error)
    fetch(`${API_ENDPOINT}/hashtags`)
      .then(res => res.json())
      .then(({ hashtags }) => setHashtags(hashtags))
      .catch(console.error)
  }, []);
  return (
    <>
      <Head>
        <title>Wisesight Test</title>
      </Head>
      <Dashboard
        topAccountsByMessages={topAccountsByMessages}
        topMessagesByEngagements={topMessagesByEngagements}
        numberOfDailyMessage={numberOfDailyMessage}
        words={words}
        hashtags={hashtags}
      />
    </>
  );
};

export default IndexPage;

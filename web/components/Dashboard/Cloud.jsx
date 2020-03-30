import React from 'react'
import dynamic from 'next/dynamic'

const fontSizeMapper = ({ value }) => Math.log2(value) * 5;

const WordCloud = dynamic(() => import("react-d3-cloud"), { ssr: false });

export const Cloud = React.memo(({ width, height, data })=> (
  <WordCloud width={width} height={height} data={data} fontSizeMapper={fontSizeMapper} />
), (prev, next) => prev.data.length === next.data.length)
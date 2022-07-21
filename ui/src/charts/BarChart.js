/** BarChart.js */
import * as d3 from "d3";
import React, { useEffect, useRef } from "react";

const BarChart = ({ data }) => {

    // const [response, isLoading] = useFetch('http://localhost:8080/v1/api/new_cases');

    const svgRef = useRef(null);

    const height = 500;
    const width = 500;
    const barWidth = (width / data.length) * 0.9;
    const marginBottom = 100;
    const chartHeight = height - marginBottom;
    const opacity = 0.65;

    const maxValue = d3.max(data, (d) => d.new_cases);
    const heightScale = d3.scaleLinear()
        .domain([0, maxValue])
        .range([0, chartHeight - 30]);

    useEffect(() => {
        const svgElement = d3.select(svgRef.current);
        svgElement.selectAll('*').remove();

        const svg = svgElement.append('svg')
            .attr('width', width)
            .attr('height', height);

        svg.selectAll('rect')
            .data(data)
            .enter()
            .append('rect')

            .attr('x', (d, i) => barWidth / 2 + i * barWidth)
            .attr('y', (d) => chartHeight - heightScale(d.new_cases))
            .attr('width', barWidth)
            .attr('height', (d) => heightScale(Math.abs(d.new_cases)))
            .attr('fill', 'darkBlue')
            .attr('opacity', opacity)

            .on('mouseover', function () {
                d3.select(this).transition()
                    .duration('50')
                    .attr('opacity', '0.30');
            })
            .on('mouseout', function (d) {
                d3.select(this).transition()
                    .duration('50')
                    .attr('opacity', opacity);
            })

        svg.selectAll('text')
            .data(data)
            .enter()
            .append('text')
            .attr('x', (d, i) => barWidth / 2 + i * barWidth)
            .text((d) => d3.format(',')(d.new_cases))
            .style('font', '10px arial')
            .attr('y', (d, i) => chartHeight - heightScale(d.new_cases) - 5)

    }, [data]);

    return <svg ref={svgRef} width={width} height={height} />;
};

export default BarChart;

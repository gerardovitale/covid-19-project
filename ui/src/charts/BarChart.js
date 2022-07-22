/** BarChart.js */
import * as d3 from "d3";
import React, { useEffect, useRef } from "react";

const BarChart = ({ data }) => {

    // const [response, isLoading] = useFetch('http://localhost:8080/v1/api/new_cases');

    const svgRef = useRef(null);

    const height = 500;
    const width = 500;
    const barWidth = (width / data.length) * 0.90;
    const marginBottom = 100;
    const chartHeight = height - marginBottom;
    const opacity = 0.65;

    const maxValue = d3.max(data, (d) => d.new_cases);
    const yScale = d3.scaleLinear()
        .domain([0, maxValue])
        .range([0, chartHeight - 30]);
    const xScale = d3.scaleLinear()
        .domain([0, data.length + 1])
        .range([0, width]);

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
            .attr('x', (d) => xScale(d.month))
            .attr('y', (d) => chartHeight - yScale(d.new_cases))
            .attr('width', barWidth)
            .attr('height', (d) => yScale(d.new_cases))
            .attr('fill', 'darkBlue')
            .attr('opacity', opacity)
            .on('mouseover', function () {
                d3.select(this).transition()
                    .duration('50')
                    .attr('opacity', '0.30');
            })
            .on('mouseout', function () {
                d3.select(this).transition()
                    .duration('50')
                    .attr('opacity', opacity);
            });

        svg.selectAll('text')
            .data(data)
            .enter()
            .append('text')
            .attr('x', (d) => xScale(d.month))
            .text((d) => d3.format(',')(d.new_cases))
            .style('font', '10px arial')
            .attr('y', (d, i) => chartHeight - yScale(d.new_cases) - 5);

        const xAxis = d3.axisBottom(xScale)
            .ticks(5)
            .tickSize(- height + marginBottom);
        const xAxisGroup = svg.append("g")
            .attr("transform", `translate(0, ${height - marginBottom})`)
            .call(xAxis);
        xAxisGroup.select(".domain").remove();
        // xAxisGroup.selectAll("rect")
        // .attr("stroke", "rgba(255, 255, 255, 0.2)");
        xAxisGroup.selectAll("text")
            .attr("opacity", 0.5)
            .attr("color", "black")
            .attr("font-size", "0.75rem");

    }, [data]);

    return <svg ref={svgRef} width={width} height={height} />;
};

export default BarChart;

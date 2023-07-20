/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: rleslie- <rleslie-@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/07/19 14:17:43 by rleslie-          #+#    #+#             */
/*   Updated: 2023/07/19 15:05:10 by rleslie-         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"os"
	"strings"
)

func	main() {
	args := os.Args[1:]
	if len(args) == 0{
		return
	}
	for i, copy := range args{
		args[i] = strings.ToUpper(copy)
	}

	fmt.Println(args)
}
